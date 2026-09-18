package entity

type Invitation struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	UserID           uint   `json:"userId" gorm:"not null"`
	BrideName        string `json:"brideName"`
	BrideDegree      string `json:"brideDegree"`
	GroomName        string `json:"groomName"`
	GroomDegree      string `json:"groomDegree"`
	GroomImgUrl      string `json:"groomImgUrl"`
	BrideImgUrl      string `json:"brideImgUrl"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	GroomDescription string `json:"groomDescription"`
	BrideDescription string `json:"brideDescription"`
}
