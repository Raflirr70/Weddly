package kafka

import "time"

const (
	TopicVisitorLog = "visitor-log"
	TopicComment    = "comment-created"
	ConsumerGroupID = "weddly-worker"
)

type VisitorLogEvent struct {
	InvitationID uint      `json:"invitation_id"`
	UserID       uint      `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	VisitedAt    time.Time `json:"visited_at"`
}

type CommentCreatedEvent struct {
	InvitationID     uint      `json:"invitation_id"`
	UserID           uint      `json:"user_id"`
	Username         string    `json:"username"`
	Comment          string    `json:"comment"`
	ConfirmAttendant bool      `json:"confirm_attendant"`
	CreatedAt        time.Time `json:"created_at"`
}
