package entity

type Story struct {
	Order       int    `json:"order"`
	StoryImgUrl string `json:"storyImgUrl"`
	Title       string `json:"title"`
	Description string `json:"description"`
}