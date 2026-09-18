package entity

import "time"

type Comment struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	UserID           uint      `json:"userId" gorm:"index"`
	Username         string    `json:"username"`
	Comment          string    `json:"comment"`
	ConfirmAttendant bool      `json:"confirmAttendant"`
	CreatedAt        time.Time `json:"createdAt"`
}
