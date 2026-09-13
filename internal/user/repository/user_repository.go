package repository

import (
	"github.com/Raflirr70/Weddly/internal/user/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

func FindByUsername(db *gorm.DB, username string) (*entity.User, error) {
	var user entity.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
