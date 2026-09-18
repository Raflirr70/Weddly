package entity

import "time"

type VisitorLog struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	InvitationID uint      `json:"invitationId" gorm:"index"`
	IPAddress    string    `json:"ipAddress"`
	UserAgent    string    `json:"userAgent"`
	VisitedAt    time.Time `json:"visitedAt"`
}

type ActivityLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"userId" gorm:"index"`
	Action    string    `json:"action"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"createdAt"`
}

type VisitorStat struct {
	InvitationID  uint      `json:"invitationId"`
	VisitCount    int64     `json:"visitCount"`
	LastVisitedAt time.Time `json:"lastVisitedAt"`
}
