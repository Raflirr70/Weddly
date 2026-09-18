package entity

type Gallery struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserID   uint   `json:"userId" gorm:"not null;index"`
	ImageUrl string `json:"imageUrl"`
}
