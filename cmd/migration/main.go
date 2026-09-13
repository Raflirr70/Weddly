package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/Raflirr70/Weddly/internal/user/entity"
	"github.com/Raflirr70/Weddly/internal/user/repository"
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
	if err := repository.Migrate(db); err != nil {
		log.Fatal("Migrate failed:", err)
	}

	// 2. Seed: isi superadmin default kalau belum ada
	username := "admin"
	password := "admin123"

	existing, err := repository.FindByUsername(db, username)
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
