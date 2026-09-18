package entity

type Opening struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	UserID        uint   `json:"userId" gorm:"not null"`
	OpeningImgUrl string `json:"openingImgUrl"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}
