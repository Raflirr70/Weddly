package main

import (
	"log"

	"github.com/gin-gonic/gin"

	authHandler "github.com/Raflirr70/Weddly/internal/auth/handler"
	authRepository "github.com/Raflirr70/Weddly/internal/auth/repository"
	authUsecase "github.com/Raflirr70/Weddly/internal/auth/usecase"
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

	repo := authRepository.NewAuthRepository(db, redisClient)
	usecase := authUsecase.NewAuthUsecase(repo, cfg.JWTSecret)
	handler := authHandler.NewAuthHandler(usecase)

	r := gin.Default()
	api := r.Group("/api/v1")
	{
		api.POST("/login", handler.Login)
	}

	log.Println("auth-service running on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}