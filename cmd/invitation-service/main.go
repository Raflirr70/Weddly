package main

import (
	"log"

	"github.com/gin-gonic/gin"

	invitationHandler "github.com/Raflirr70/Weddly/internal/invitation/handler"
	invitationRepository "github.com/Raflirr70/Weddly/internal/invitation/repository"
	invitationRoutes "github.com/Raflirr70/Weddly/internal/invitation/routes"
	invitationUsecase "github.com/Raflirr70/Weddly/internal/invitation/usecase"
	"github.com/Raflirr70/Weddly/pkg/config"
	"github.com/Raflirr70/Weddly/pkg/database"
	"github.com/Raflirr70/Weddly/pkg/kafka"
)

// @title         Weddly Invitation Service API
// @version       1.0
// @description   Kelola data undangan (cover, hero, opening, invitation, event, gallery, story, gift) + endpoint publik tamu.
// @host          localhost:8080
// @BasePath      /api/v1
// @securityDefinitions.apikey BearerAuth
// @in             header
// @name           Authorization
func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("PostgreSQL connect failed:", err)
	}

	redisClient, err := database.ConnectRedis(cfg)
	if err != nil {
		log.Fatal("Redis connect failed:", err)
	}

	repo := invitationRepository.NewInvitationRepository(db)
	usecase := invitationUsecase.NewInvitationUsecase(repo)
	producer := kafka.NewProducer(cfg.KafkaBroker)
	handler := invitationHandler.NewInvitationHandler(usecase, producer)

	r := gin.Default()
	invitationRoutes.NewInvitationRoutes(handler, redisClient, cfg.JWTSecret).Register(r.Group("/api/v1"))

	log.Println("invitation-service running on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
