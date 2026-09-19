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

type EventEntryRequest struct {
	Order        int       `json:"order"`
	Title        string    `json:"title" binding:"required"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}

type EventRequest struct {
	Title  string              `json:"title"`
	Events []EventEntryRequest `json:"events"`
}

type GalleryEntryRequest struct {
	Order    int    `json:"order"`
	ImageUrl string `json:"imageUrl" binding:"required"`
}

type GalleryRequest struct {
	Galleries []GalleryEntryRequest `json:"galleries"`
}

type StoryEntryRequest struct {
	Order       int    `json:"order"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type StoryRequest struct {
	StoryImgUrl string              `json:"storyImgUrl"`
	Stories     []StoryEntryRequest `json:"stories"`
}

type GiftEntryRequest struct {
	Order           int    `json:"order"`
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}

type GiftRequest struct {
	Gifts []GiftEntryRequest `json:"gifts"`
}
