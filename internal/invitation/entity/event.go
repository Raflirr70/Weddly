package entity

import "time"

type Event struct {
	Order        int       `json:"order"`
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}