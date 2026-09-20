package main

import (
	"log"

	"github.com/gin-gonic/gin"

	authHandler "github.com/Raflirr70/Weddly/internal/auth/handler"
	authRepository "github.com/Raflirr70/Weddly/internal/auth/repository"
	authRoutes "github.com/Raflirr70/Weddly/internal/auth/routes"
	authUsecase "github.com/Raflirr70/Weddly/internal/auth/usecase"
	notificationHandler "github.com/Raflirr70/Weddly/internal/notification/handler"
	notificationRepository "github.com/Raflirr70/Weddly/internal/notification/repository"
	notificationRoutes "github.com/Raflirr70/Weddly/internal/notification/routes"
	notificationUsecase "github.com/Raflirr70/Weddly/internal/notification/usecase"
	userHandler "github.com/Raflirr70/Weddly/internal/user/handler"
	userRepository "github.com/Raflirr70/Weddly/internal/user/repository"
	userRoutes "github.com/Raflirr70/Weddly/internal/user/routes"
	userUsecase "github.com/Raflirr70/Weddly/internal/user/usecase"
	"github.com/Raflirr70/Weddly/pkg/config"
	"github.com/Raflirr70/Weddly/pkg/database"
)

// @title         Weddly Auth Service API
// @version       1.0
// @description   Auth, kelola user & monitoring (login, user, logs, visitors).
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

	repo := authRepository.NewAuthRepository(db, redisClient)
	usecase := authUsecase.NewAuthUsecase(repo, cfg.JWTSecret)
	handler := authHandler.NewAuthHandler(usecase)

	userRepo := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler := userHandler.NewUserHandler(userUsecase)

	notificationRepo := notificationRepository.NewNotificationRepository(db)
	notificationUsecase := notificationUsecase.NewNotificationUsecase(notificationRepo)
	notificationHandler := notificationHandler.NewNotificationHandler(notificationUsecase)

	r := gin.Default()
	authRoutes.NewAuthRoutes(handler, redisClient, cfg.JWTSecret).Register(r.Group("/api/v1"))
	userRoutes.NewUserRoutes(userHandler, redisClient, cfg.JWTSecret).Register(r.Group("/api/v1"))
	notificationRoutes.NewNotificationRoutes(notificationHandler, redisClient, cfg.JWTSecret).Register(r.Group("/api/v1"))

	log.Println("auth-service running on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
