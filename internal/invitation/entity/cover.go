package entity

type Cover struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"userId" gorm:"not null"`
	CoverUrl string `json:"coverUrl"`
}
