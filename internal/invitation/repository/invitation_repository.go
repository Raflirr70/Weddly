package repository

import (
	"gorm.io/gorm"

	"github.com/Raflirr70/Weddly/internal/invitation/entity"
)

type InvitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository(db *gorm.DB) *InvitationRepository {
	return &InvitationRepository{db: db}
}

func (r *InvitationRepository) Migrate() error {
	return r.db.AutoMigrate(
		&entity.Invitation{},
		&entity.Comment{},
	)
}

// ---------- Invitation ----------

func (r *InvitationRepository) FindInvitationByID(id uint) (*entity.Invitation, error) {
	var inv entity.Invitation
	if err := r.db.First(&inv, id).Error; err != nil {
		return nil, err
	}
	return &inv, nil
}

func (r *InvitationRepository) FindInvitationByUserID(userID uint) (*entity.Invitation, error) {
	var inv entity.Invitation
	if err := r.db.Where("user_id = ?", userID).First(&inv).Error; err != nil {
		return nil, err
	}
	return &inv, nil
}

func (r *InvitationRepository) CreateInvitation(inv *entity.Invitation) error {
	return r.db.Create(inv).Error
}

func (r *InvitationRepository) UpdateInvitation(inv *entity.Invitation) error {
	return r.db.Save(inv).Error
}

func (r *InvitationRepository) DeleteInvitation(id uint) error {
	return r.db.Delete(&entity.Invitation{}, id).Error
}

// ---------- Comment ----------

func (r *InvitationRepository) CreateComment(c *entity.Comment) error {
	return r.db.Create(c).Error
}

func (r *InvitationRepository) FindCommentsByUserID(userID uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}