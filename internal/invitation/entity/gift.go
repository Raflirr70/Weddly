package entity

type Gift struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	UserID          uint   `json:"userId" gorm:"not null"`
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}
