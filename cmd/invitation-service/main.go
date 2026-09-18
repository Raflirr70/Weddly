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
)

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
	handler := invitationHandler.NewInvitationHandler(usecase)

	r := gin.Default()
	invitationRoutes.NewInvitationRoutes(handler, redisClient, cfg.JWTSecret).Register(r.Group("/api/v1"))

	log.Println("invitation-service running on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
