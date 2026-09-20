package entity

import "time"

type Event struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	InvitationID uint      `json:"invitationId" gorm:"index"`
	Order        int       `json:"order" gorm:"column:sort_order"`
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}