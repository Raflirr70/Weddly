package entity

type Story struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"userId" gorm:"not null"`
	StoryImgUrl string `json:"storyImgUrl"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
