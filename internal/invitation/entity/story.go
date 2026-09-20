package entity

type Story struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	InvitationID uint   `json:"invitationId" gorm:"index"`
	Order        int    `json:"order" gorm:"column:sort_order"`
	StoryImgUrl  string `json:"storyImgUrl"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}
