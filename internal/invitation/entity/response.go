package entity

import "time"

type InvitationDetailResponse struct {
	ID             uint                      `json:"id"`
	UserID         uint                      `json:"-"`
	BrideName      string                    `json:"brideName"`
	BrideDegree    string                    `json:"brideDegree"`
	GroomName      string                    `json:"groomName"`
	GroomDegree    string                    `json:"groomDegree"`
	CoverUrl       string                    `json:"coverUrl"`
	HeroImgUrl     string                    `json:"heroImgUrl"`
	SectionOrder   []string                  `json:"sectionOrder"`
	SectionEnabled map[string]bool           `json:"sectionEnabled"`
	Opening        OpeningResponse           `json:"opening"`
	Invitation     InvitationSectionResponse `json:"invitation"`
	EventTime      []EventResponse           `json:"eventTime"`
	Galery         []string                  `json:"galery"`
	StorySection   StorySectionResponse      `json:"storySection"`
	Gift           []GiftResponse            `json:"gift"`
}

type OpeningResponse struct {
	OpeningImageUrl string `json:"openingImageUrl"`
	Title           string `json:"title"`
	Description     string `json:"description"`
}

type InvitationSectionResponse struct {
	GroomImgUrl      string `json:"groomImgUrl"`
	BrideImgUrl      string `json:"brideImgUrl"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	GroomDescription string `json:"groomDescription"`
	BrideDescription string `json:"brideDescription"`
}

type EventResponse struct {
	Title        string    `json:"title"`
	Location     string    `json:"location"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	LocationLink string    `json:"locationLink"`
}

type StorySectionResponse struct {
	StoryImgUrl string       `json:"storyImgUrl"`
	Story       []StoryEntry `json:"story"`
}

type StoryEntry struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GiftResponse struct {
	Provider        string `json:"provider"`
	ProviderAccount string `json:"providerAccount"`
	Type            string `json:"type"`
	No              string `json:"no"`
}

type CommentResponse struct {
	InvitationID     uint      `json:"-"`
	UserID           uint      `json:"-"`
	Username         string    `json:"username"`
	Comment          string    `json:"comment"`
	ConfirmAttendant bool      `json:"confirmAttendant"`
	CreatedAt        time.Time `json:"createdAt"`
}

type CommentRequest struct {
	Username         string `json:"username" binding:"required"`
	Comment          string `json:"comment" binding:"required"`
	ConfirmAttendant bool   `json:"confirmAttendant"`
}
