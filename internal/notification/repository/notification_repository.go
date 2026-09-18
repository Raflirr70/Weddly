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

func (r *NotificationRepository) FindActivityLogs() ([]entity.ActivityLog, error) {
	var logs []entity.ActivityLog
	if err := r.db.Order("created_at DESC").Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *NotificationRepository) CountVisitorsByInvitation() ([]entity.VisitorStat, error) {
	var stats []entity.VisitorStat
	if err := r.db.Model(&entity.VisitorLog{}).
		Select("invitation_id, count(*) as visit_count, max(visited_at) as last_visited_at").
		Group("invitation_id").
		Order("visit_count DESC").
		Scan(&stats).Error; err != nil {
		return nil, err
	}
	return stats, nil
}
