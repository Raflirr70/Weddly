package entity

type Hero struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	InvitationID uint   `json:"invitationId" gorm:"index"`
	HeroImgUrl   string `json:"heroImgUrl"`
}