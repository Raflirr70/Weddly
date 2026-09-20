package repository

import (
	"errors"

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
		&entity.Cover{},
		&entity.Hero{},
		&entity.Opening{},
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
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, model := range []interface{}{
			&entity.Cover{}, &entity.Hero{}, &entity.Opening{}, &entity.Event{},
			&entity.Gallery{}, &entity.Story{}, &entity.Gift{},
		} {
			if err := tx.Where("invitation_id = ?", id).Delete(model).Error; err != nil {
				return err
			}
		}
		return tx.Delete(&entity.Invitation{}, id).Error
	})
}

// ---------- Cover / Hero / Opening (satu baris per undangan) ----------

func (r *InvitationRepository) UpsertCover(invitationID uint, c *entity.Cover) error {
	c.InvitationID = invitationID
	var existing entity.Cover
	if err := r.db.Where("invitation_id = ?", invitationID).First(&existing).Error; err == nil {
		c.ID = existing.ID
		return r.db.Save(c).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(c).Error
	} else {
		return err
	}
}

func (r *InvitationRepository) UpsertHero(invitationID uint, h *entity.Hero) error {
	h.InvitationID = invitationID
	var existing entity.Hero
	if err := r.db.Where("invitation_id = ?", invitationID).First(&existing).Error; err == nil {
		h.ID = existing.ID
		return r.db.Save(h).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(h).Error
	} else {
		return err
	}
}

func (r *InvitationRepository) UpsertOpening(invitationID uint, o *entity.Opening) error {
	o.InvitationID = invitationID
	var existing entity.Opening
	if err := r.db.Where("invitation_id = ?", invitationID).First(&existing).Error; err == nil {
		o.ID = existing.ID
		return r.db.Save(o).Error
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.db.Create(o).Error
	} else {
		return err
	}
}

func (r *InvitationRepository) GetCover(invitationID uint) (*entity.Cover, error) {
	var c entity.Cover
	if err := r.db.Where("invitation_id = ?", invitationID).First(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *InvitationRepository) GetHero(invitationID uint) (*entity.Hero, error) {
	var h entity.Hero
	if err := r.db.Where("invitation_id = ?", invitationID).First(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}

func (r *InvitationRepository) GetOpening(invitationID uint) (*entity.Opening, error) {
	var o entity.Opening
	if err := r.db.Where("invitation_id = ?", invitationID).First(&o).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *InvitationRepository) DeleteCover(invitationID uint) error {
	return r.db.Where("invitation_id = ?", invitationID).Delete(&entity.Cover{}).Error
}

func (r *InvitationRepository) DeleteHero(invitationID uint) error {
	return r.db.Where("invitation_id = ?", invitationID).Delete(&entity.Hero{}).Error
}

func (r *InvitationRepository) DeleteOpening(invitationID uint) error {
	return r.db.Where("invitation_id = ?", invitationID).Delete(&entity.Opening{}).Error
}

// ---------- Event / Gallery / Story / Gift (batch replace) ----------

func (r *InvitationRepository) ReplaceEvents(invitationID uint, items []entity.Event) error {
	return r.replaceCollection(invitationID, items, func() interface{} { return &entity.Event{} })
}

func (r *InvitationRepository) ReplaceGalleries(invitationID uint, items []entity.Gallery) error {
	return r.replaceCollection(invitationID, items, func() interface{} { return &entity.Gallery{} })
}

func (r *InvitationRepository) ReplaceStories(invitationID uint, items []entity.Story) error {
	return r.replaceCollection(invitationID, items, func() interface{} { return &entity.Story{} })
}

func (r *InvitationRepository) ReplaceGifts(invitationID uint, items []entity.Gift) error {
	return r.replaceCollection(invitationID, items, func() interface{} { return &entity.Gift{} })
}

func (r *InvitationRepository) replaceCollection(invitationID uint, items interface{}, model func() interface{}) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("invitation_id = ?", invitationID).Delete(model()).Error; err != nil {
			return err
		}
		return tx.Create(items).Error
	})
}

func (r *InvitationRepository) GetEvents(invitationID uint) ([]entity.Event, error) {
	var items []entity.Event
	if err := r.db.Where("invitation_id = ?", invitationID).Order("sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InvitationRepository) GetGalleries(invitationID uint) ([]entity.Gallery, error) {
	var items []entity.Gallery
	if err := r.db.Where("invitation_id = ?", invitationID).Order("sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InvitationRepository) GetStories(invitationID uint) ([]entity.Story, error) {
	var items []entity.Story
	if err := r.db.Where("invitation_id = ?", invitationID).Order("sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *InvitationRepository) GetGifts(invitationID uint) ([]entity.Gift, error) {
	var items []entity.Gift
	if err := r.db.Where("invitation_id = ?", invitationID).Order("sort_order ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
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