package usecase

import (
	"github.com/Raflirr70/Weddly/internal/notification/entity"
	"github.com/Raflirr70/Weddly/internal/notification/repository"
	"github.com/Raflirr70/Weddly/pkg/apperror"
)

type NotificationUsecase struct {
	repo *repository.NotificationRepository
}

func NewNotificationUsecase(repo *repository.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{repo: repo}
}

func (u *NotificationUsecase) ListLogs() ([]entity.ActivityLog, *apperror.AppError) {
	logs, err := u.repo.FindActivityLogs()
	if err != nil {
		return nil, apperror.Internal("Failed to list activity logs")
	}
	return logs, nil
}

func (u *NotificationUsecase) VisitorStats() ([]entity.VisitorStat, *apperror.AppError) {
	stats, err := u.repo.CountVisitorsByInvitation()
	if err != nil {
		return nil, apperror.Internal("Failed to get visitor stats")
	}
	return stats, nil
}
