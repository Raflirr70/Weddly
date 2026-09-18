package entity

import "time"

type CoverRequest struct {
	CoverUrl string `json:"coverUrl" binding:"required"`
}

type HeroRequest struct {
	HeroImgUrl string `json:"heroImgUrl" binding:"required"`
}

type OpeningRequest struct {
	OpeningImgUrl string `json:"openingImgUrl"`
	Title         string `json:"title"`
	Description   string `json:"description"`
}

type InvitationRequest struct {
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

type EventRequest struct {
	Title        string    `json:"title" binding:"required"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}

type GalleryRequest struct {
	ImageUrl string `json:"imageUrl" binding:"required"`
}

type StoryRequest struct {
	StoryImgUrl string `json:"storyImgUrl"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GiftRequest struct {
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}
