package entity

type Opening struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	InvitationID  uint   `json:"invitationId" gorm:"index"`
	OpeningImgUrl string `json:"openingImgUrl"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}