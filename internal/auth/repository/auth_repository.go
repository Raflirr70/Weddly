package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/Raflirr70/Weddly/internal/user/entity"
)

type AuthRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewAuthRepository(db *gorm.DB, redis *redis.Client) *AuthRepository {
	return &AuthRepository{db: db, redis: redis}
}

func (r *AuthRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) SaveToken(userID uint, token string, ttl time.Duration) error {
	key := "auth:token:" + strconv.FormatUint(uint64(userID), 10)
	return r.redis.Set(context.Background(), key, token, ttl).Err()
}