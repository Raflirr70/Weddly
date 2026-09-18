package repository

import (
	"gorm.io/gorm"

	"github.com/Raflirr70/Weddly/internal/notification/entity"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Migrate() error {
	return r.db.AutoMigrate(&entity.VisitorLog{}, &entity.ActivityLog{})
}

func (r *NotificationRepository) CreateVisitorLog(log *entity.VisitorLog) error {
	return r.db.Create(log).Error
}

func (r *NotificationRepository) CreateActivityLog(log *entity.ActivityLog) error {
	return r.db.Create(log).Error
}
