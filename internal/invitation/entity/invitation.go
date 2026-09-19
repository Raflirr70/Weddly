package entity

type Invitation struct {
	ID               uint               `json:"id" gorm:"primaryKey"`
	UserID           uint               `json:"userId" gorm:"not null;index:idx_invitations_user,unique"`
	BrideName        string             `json:"brideName"`
	BrideDegree      string             `json:"brideDegree"`
	GroomName        string             `json:"groomName"`
	GroomDegree      string             `json:"groomDegree"`
	GroomImgUrl      string             `json:"groomImgUrl"`
	BrideImgUrl      string             `json:"brideImgUrl"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	GroomDescription string             `json:"groomDescription"`
	BrideDescription string             `json:"brideDescription"`
	Sections         InvitationSections `json:"-" gorm:"serializer:json;type:jsonb"`
}

type InvitationSections struct {
	CoverUrl       string          `json:"coverUrl"`
	HeroImgUrl     string          `json:"heroImgUrl"`
	Opening        Opening         `json:"opening"`
	Events         []Event         `json:"events"`
	Galleries      []Gallery       `json:"galleries"`
	Stories        []Story         `json:"stories"`
	Gifts          []Gift          `json:"gifts"`
	SectionOrder   []string        `json:"sectionOrder"`
	SectionEnabled map[string]bool `json:"sectionEnabled"`
}