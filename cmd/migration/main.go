package main

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	invitationEntity "github.com/Raflirr70/Weddly/internal/invitation/entity"
	invitationRepository "github.com/Raflirr70/Weddly/internal/invitation/repository"
	notificationRepository "github.com/Raflirr70/Weddly/internal/notification/repository"
	"github.com/Raflirr70/Weddly/internal/user/entity"
	userRepository "github.com/Raflirr70/Weddly/internal/user/repository"
	"github.com/Raflirr70/Weddly/pkg/config"
	"github.com/Raflirr70/Weddly/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("PostgreSQL connect failed:", err)
	}

	// 1. Migrate: buat tabel users
	userRepo := userRepository.NewUserRepository(db)
	if err := userRepo.Migrate(); err != nil {
		log.Fatal("Migrate failed:", err)
	}

	// 1b. Migrate: buat tabel undangan + section (covers, heros, openings, events, galleries, stories, gifts, comments)
	invitationRepo := invitationRepository.NewInvitationRepository(db)
	if err := invitationRepo.Migrate(); err != nil {
		log.Fatal("Invitation migrate failed:", err)
	}

	// 1c. Migrate: buat tabel log & aktivitas
	notificationRepo := notificationRepository.NewNotificationRepository(db)
	if err := notificationRepo.Migrate(); err != nil {
		log.Fatal("Notification migrate failed:", err)
	}

	// 1d. Backfill: pastikan tiap user punya undangan (idempotent)
	users, err := userRepo.FindAll()
	if err != nil {
		log.Fatal("Backfill list users failed:", err)
	}
	for _, usr := range users {
		if _, err := invitationRepo.FindInvitationByUserID(usr.ID); errors.Is(err, gorm.ErrRecordNotFound) {
			if err := invitationRepo.CreateInvitation(&invitationEntity.Invitation{UserID: usr.ID}); err != nil {
				log.Fatalf("Backfill invitation user %d failed: %v", usr.ID, err)
			}
			log.Printf("Backfill: invitation dibuat untuk user %d", usr.ID)
		}
	}

	// 2. Seed: isi superadmin default kalau belum ada
	username := "admin"
	password := "admin123"

	existing, err := userRepo.FindByUsername(username)
	if err == nil && existing != nil {
		log.Println("Superadmin sudah ada, tidak perlu seed ulang.")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Hash password failed:", err)
	}

	user := entity.User{
		Username: username,
		Password: string(hash),
		Role:     "superadmin",
		Status:   true,
	}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Seed failed:", err)
	}

	log.Printf("Superadmin berhasil dibuat: %s", username)
}