package entity

type Gallery struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	InvitationID uint   `json:"invitationId" gorm:"index"`
	Order        int    `json:"order" gorm:"column:sort_order"`
	ImageUrl     string `json:"imageUrl"`
}