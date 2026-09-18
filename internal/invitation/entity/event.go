package entity

import "time"

type Event struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"userId" gorm:"not null"`
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}
