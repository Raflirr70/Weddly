package entity

type Hero struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"userId" gorm:"not null;index"`
	HeroImgUrl string `json:"heroImgUrl"`
}
