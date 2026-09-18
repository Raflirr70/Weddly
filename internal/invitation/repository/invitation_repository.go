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
		&entity.Cover{},
		&entity.Hero{},
		&entity.Opening{},
		&entity.Invitation{},
		&entity.Event{},
		&entity.Gallery{},
		&entity.Story{},
		&entity.Gift{},
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

// ---------- Cover ----------

func (r *InvitationRepository) CreateCover(c *entity.Cover) error {
	return r.db.Create(c).Error
}

func (r *InvitationRepository) UpdateCover(c *entity.Cover) error {
	return r.db.Save(c).Error
}

func (r *InvitationRepository) DeleteCover(id uint) error {
	return r.db.Delete(&entity.Cover{}, id).Error
}

func (r *InvitationRepository) FindCoverByID(id uint) (*entity.Cover, error) {
	var c entity.Cover
	if err := r.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *InvitationRepository) FindCoverByUserID(userID uint) (*entity.Cover, error) {
	var c entity.Cover
	if err := r.db.Where("user_id = ?", userID).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// ---------- Hero ----------

func (r *InvitationRepository) CreateHero(h *entity.Hero) error {
	return r.db.Create(h).Error
}

func (r *InvitationRepository) UpdateHero(h *entity.Hero) error {
	return r.db.Save(h).Error
}

func (r *InvitationRepository) DeleteHero(id uint) error {
	return r.db.Delete(&entity.Hero{}, id).Error
}

func (r *InvitationRepository) FindHeroByID(id uint) (*entity.Hero, error) {
	var h entity.Hero
	if err := r.db.First(&h, id).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *InvitationRepository) FindHeroByUserID(userID uint) (*entity.Hero, error) {
	var h entity.Hero
	if err := r.db.Where("user_id = ?", userID).First(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

// ---------- Opening ----------

func (r *InvitationRepository) CreateOpening(o *entity.Opening) error {
	return r.db.Create(o).Error
}

func (r *InvitationRepository) UpdateOpening(o *entity.Opening) error {
	return r.db.Save(o).Error
}

func (r *InvitationRepository) DeleteOpening(id uint) error {
	return r.db.Delete(&entity.Opening{}, id).Error
}

func (r *InvitationRepository) FindOpeningByID(id uint) (*entity.Opening, error) {
	var o entity.Opening
	if err := r.db.First(&o, id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *InvitationRepository) FindOpeningByUserID(userID uint) (*entity.Opening, error) {
	var o entity.Opening
	if err := r.db.Where("user_id = ?", userID).First(&o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

// ---------- Event ----------

func (r *InvitationRepository) CreateEvent(e *entity.Event) error {
	return r.db.Create(e).Error
}

func (r *InvitationRepository) UpdateEvent(e *entity.Event) error {
	return r.db.Save(e).Error
}

func (r *InvitationRepository) DeleteEvent(id uint) error {
	return r.db.Delete(&entity.Event{}, id).Error
}

func (r *InvitationRepository) FindEventByID(id uint) (*entity.Event, error) {
	var e entity.Event
	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *InvitationRepository) FindEventsByUserID(userID uint) ([]entity.Event, error) {
	var events []entity.Event
	if err := r.db.Where("user_id = ?", userID).Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

// ---------- Gallery ----------

func (r *InvitationRepository) CreateGallery(g *entity.Gallery) error {
	return r.db.Create(g).Error
}

func (r *InvitationRepository) DeleteGallery(id uint) error {
	return r.db.Delete(&entity.Gallery{}, id).Error
}

func (r *InvitationRepository) FindGalleryByID(id uint) (*entity.Gallery, error) {
	var g entity.Gallery
	if err := r.db.First(&g, id).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *InvitationRepository) FindGalleriesByUserID(userID uint) ([]entity.Gallery, error) {
	var galleries []entity.Gallery
	if err := r.db.Where("user_id = ?", userID).Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}

// ---------- Story ----------

func (r *InvitationRepository) CreateStory(s *entity.Story) error {
	return r.db.Create(s).Error
}

func (r *InvitationRepository) UpdateStory(s *entity.Story) error {
	return r.db.Save(s).Error
}

func (r *InvitationRepository) DeleteStory(id uint) error {
	return r.db.Delete(&entity.Story{}, id).Error
}

func (r *InvitationRepository) FindStoryByID(id uint) (*entity.Story, error) {
	var s entity.Story
	if err := r.db.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *InvitationRepository) FindStoriesByUserID(userID uint) ([]entity.Story, error) {
	var stories []entity.Story
	if err := r.db.Where("user_id = ?", userID).Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

// ---------- Gift ----------

func (r *InvitationRepository) CreateGift(g *entity.Gift) error {
	return r.db.Create(g).Error
}

func (r *InvitationRepository) UpdateGift(g *entity.Gift) error {
	return r.db.Save(g).Error
}

func (r *InvitationRepository) DeleteGift(id uint) error {
	return r.db.Delete(&entity.Gift{}, id).Error
}

func (r *InvitationRepository) FindGiftByID(id uint) (*entity.Gift, error) {
	var g entity.Gift
	if err := r.db.First(&g, id).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *InvitationRepository) FindGiftsByUserID(userID uint) ([]entity.Gift, error) {
	var gifts []entity.Gift
	if err := r.db.Where("user_id = ?", userID).Find(&gifts).Error; err != nil {
		return nil, err
	}
	return gifts, nil
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
