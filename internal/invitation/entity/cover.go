package entity

type Cover struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	InvitationID uint   `json:"invitationId" gorm:"index"`
	CoverUrl     string `json:"coverUrl"`
}