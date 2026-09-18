package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	notificationEntity "github.com/Raflirr70/Weddly/internal/notification/entity"
	notificationRepository "github.com/Raflirr70/Weddly/internal/notification/repository"
	"github.com/Raflirr70/Weddly/pkg/config"
	"github.com/Raflirr70/Weddly/pkg/database"
	"github.com/Raflirr70/Weddly/pkg/kafka"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("PostgreSQL connect failed:", err)
	}

	repo := notificationRepository.NewNotificationRepository(db)
	if err := repo.Migrate(); err != nil {
		log.Fatal("Migrate notification failed:", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	consumer := kafka.NewConsumer(cfg.KafkaBroker, []string{kafka.TopicVisitorLog, kafka.TopicComment}, kafka.ConsumerGroupID)

	handle := func(topic string, data []byte) error {
		switch topic {
		case kafka.TopicVisitorLog:
			var ev kafka.VisitorLogEvent
			if err := json.Unmarshal(data, &ev); err != nil {
				return err
			}
			return repo.CreateVisitorLog(&notificationEntity.VisitorLog{
				InvitationID: ev.InvitationID,
				IPAddress:    ev.IPAddress,
				UserAgent:    ev.UserAgent,
				VisitedAt:    ev.VisitedAt,
			})
		case kafka.TopicComment:
			var ev kafka.CommentCreatedEvent
			if err := json.Unmarshal(data, &ev); err != nil {
				return err
			}
			detail, _ := json.Marshal(ev)
			return repo.CreateActivityLog(&notificationEntity.ActivityLog{
				UserID:    ev.UserID,
				Action:    "comment_created",
				Detail:    string(detail),
				CreatedAt: ev.CreatedAt,
			})
		}
		return nil
	}

	log.Printf("worker listening topics: %s, %s", kafka.TopicVisitorLog, kafka.TopicComment)
	consumer.Start(ctx, handle)
}
