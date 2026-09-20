package entity

type Gift struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	InvitationID    uint   `json:"invitationId" gorm:"index"`
	Order           int    `json:"order" gorm:"column:sort_order"`
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}