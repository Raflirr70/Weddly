package usecase

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Raflirr70/Weddly/internal/invitation/entity"
	"github.com/Raflirr70/Weddly/internal/invitation/repository"
	"github.com/Raflirr70/Weddly/pkg/apperror"
)

type InvitationUsecase struct {
	repo *repository.InvitationRepository
}

func NewInvitationUsecase(repo *repository.InvitationRepository) *InvitationUsecase {
	return &InvitationUsecase{repo: repo}
}

func mapRepoErr(err error, notFoundMsg string) *apperror.AppError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.NotFound(notFoundMsg)
	}
	return apperror.Internal("Database error")
}

func forbidden() *apperror.AppError {
	return apperror.Forbidden("Forbidden")
}

// ---------- Cover ----------

func (u *InvitationUsecase) CreateCover(userID uint, req entity.CoverRequest) (*entity.Cover, *apperror.AppError) {
	cover := entity.Cover{UserID: userID, CoverUrl: req.CoverUrl}
	if err := u.repo.CreateCover(&cover); err != nil {
		return nil, apperror.Internal("Failed to create cover")
	}
	return &cover, nil
}

func (u *InvitationUsecase) DeleteCover(userID, id uint) *apperror.AppError {
	cover, err := u.repo.FindCoverByID(id)
	if err != nil {
		return mapRepoErr(err, "Cover Not Found")
	}
	if cover.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteCover(id); err != nil {
		return apperror.Internal("Failed to delete cover")
	}
	return nil
}

// ---------- Hero ----------

func (u *InvitationUsecase) CreateHero(userID uint, req entity.HeroRequest) (*entity.Hero, *apperror.AppError) {
	hero := entity.Hero{UserID: userID, HeroImgUrl: req.HeroImgUrl}
	if err := u.repo.CreateHero(&hero); err != nil {
		return nil, apperror.Internal("Failed to create hero")
	}
	return &hero, nil
}

func (u *InvitationUsecase) DeleteHero(userID, id uint) *apperror.AppError {
	hero, err := u.repo.FindHeroByID(id)
	if err != nil {
		return mapRepoErr(err, "Hero Not Found")
	}
	if hero.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteHero(id); err != nil {
		return apperror.Internal("Failed to delete hero")
	}
	return nil
}

// ---------- Opening ----------

func (u *InvitationUsecase) CreateOpening(userID uint, req entity.OpeningRequest) (*entity.Opening, *apperror.AppError) {
	opening := entity.Opening{
		UserID:        userID,
		OpeningImgUrl: req.OpeningImgUrl,
		Title:         req.Title,
		Description:   req.Description,
	}
	if err := u.repo.CreateOpening(&opening); err != nil {
		return nil, apperror.Internal("Failed to create opening")
	}
	return &opening, nil
}

func (u *InvitationUsecase) UpdateOpening(userID, id uint, req entity.OpeningRequest) (*entity.Opening, *apperror.AppError) {
	opening, err := u.repo.FindOpeningByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Opening Not Found")
	}
	if opening.UserID != userID {
		return nil, forbidden()
	}
	opening.OpeningImgUrl = req.OpeningImgUrl
	opening.Title = req.Title
	opening.Description = req.Description
	if err := u.repo.UpdateOpening(opening); err != nil {
		return nil, apperror.Internal("Failed to update opening")
	}
	return opening, nil
}

func (u *InvitationUsecase) DeleteOpening(userID, id uint) *apperror.AppError {
	opening, err := u.repo.FindOpeningByID(id)
	if err != nil {
		return mapRepoErr(err, "Opening Not Found")
	}
	if opening.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteOpening(id); err != nil {
		return apperror.Internal("Failed to delete opening")
	}
	return nil
}

// ---------- Invitation ----------

func (u *InvitationUsecase) CreateInvitation(userID uint, req entity.InvitationRequest) (*entity.Invitation, *apperror.AppError) {
	inv := entity.Invitation{
		UserID:           userID,
		BrideName:        req.BrideName,
		BrideDegree:      req.BrideDegree,
		GroomName:        req.GroomName,
		GroomDegree:      req.GroomDegree,
		GroomImgUrl:      req.GroomImgUrl,
		BrideImgUrl:      req.BrideImgUrl,
		Title:            req.Title,
		Description:      req.Description,
		GroomDescription: req.GroomDescription,
		BrideDescription: req.BrideDescription,
	}
	if err := u.repo.CreateInvitation(&inv); err != nil {
		return nil, apperror.Internal("Failed to create invitation")
	}
	return &inv, nil
}

func (u *InvitationUsecase) UpdateInvitation(userID, id uint, req entity.InvitationRequest) (*entity.Invitation, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}
	if inv.UserID != userID {
		return nil, forbidden()
	}
	inv.BrideName = req.BrideName
	inv.BrideDegree = req.BrideDegree
	inv.GroomName = req.GroomName
	inv.GroomDegree = req.GroomDegree
	inv.GroomImgUrl = req.GroomImgUrl
	inv.BrideImgUrl = req.BrideImgUrl
	inv.Title = req.Title
	inv.Description = req.Description
	inv.GroomDescription = req.GroomDescription
	inv.BrideDescription = req.BrideDescription
	if err := u.repo.UpdateInvitation(inv); err != nil {
		return nil, apperror.Internal("Failed to update invitation")
	}
	return inv, nil
}

func (u *InvitationUsecase) DeleteInvitation(userID, id uint) *apperror.AppError {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return mapRepoErr(err, "Invitation Not Found")
	}
	if inv.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteInvitation(id); err != nil {
		return apperror.Internal("Failed to delete invitation")
	}
	return nil
}

// ---------- Event ----------

func (u *InvitationUsecase) CreateEvent(userID uint, req entity.EventRequest) (*entity.Event, *apperror.AppError) {
	event := entity.Event{
		UserID:       userID,
		Title:        req.Title,
		Location:     req.Location,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		LocationLink: req.LocationLink,
	}
	if err := u.repo.CreateEvent(&event); err != nil {
		return nil, apperror.Internal("Failed to create event")
	}
	return &event, nil
}

func (u *InvitationUsecase) UpdateEvent(userID, id uint, req entity.EventRequest) (*entity.Event, *apperror.AppError) {
	event, err := u.repo.FindEventByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Event Not Found")
	}
	if event.UserID != userID {
		return nil, forbidden()
	}
	event.Title = req.Title
	event.Location = req.Location
	event.StartDate = req.StartDate
	event.EndDate = req.EndDate
	event.LocationLink = req.LocationLink
	if err := u.repo.UpdateEvent(event); err != nil {
		return nil, apperror.Internal("Failed to update event")
	}
	return event, nil
}

func (u *InvitationUsecase) DeleteEvent(userID, id uint) *apperror.AppError {
	event, err := u.repo.FindEventByID(id)
	if err != nil {
		return mapRepoErr(err, "Event Not Found")
	}
	if event.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteEvent(id); err != nil {
		return apperror.Internal("Failed to delete event")
	}
	return nil
}

// ---------- Gallery ----------

func (u *InvitationUsecase) CreateGallery(userID uint, req entity.GalleryRequest) (*entity.Gallery, *apperror.AppError) {
	gallery := entity.Gallery{UserID: userID, ImageUrl: req.ImageUrl}
	if err := u.repo.CreateGallery(&gallery); err != nil {
		return nil, apperror.Internal("Failed to create gallery")
	}
	return &gallery, nil
}

func (u *InvitationUsecase) DeleteGallery(userID, id uint) *apperror.AppError {
	gallery, err := u.repo.FindGalleryByID(id)
	if err != nil {
		return mapRepoErr(err, "Gallery Not Found")
	}
	if gallery.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteGallery(id); err != nil {
		return apperror.Internal("Failed to delete gallery")
	}
	return nil
}

// ---------- Story ----------

func (u *InvitationUsecase) CreateStory(userID uint, req entity.StoryRequest) (*entity.Story, *apperror.AppError) {
	story := entity.Story{
		UserID:      userID,
		StoryImgUrl: req.StoryImgUrl,
		Title:       req.Title,
		Description: req.Description,
	}
	if err := u.repo.CreateStory(&story); err != nil {
		return nil, apperror.Internal("Failed to create story")
	}
	return &story, nil
}

func (u *InvitationUsecase) UpdateStory(userID, id uint, req entity.StoryRequest) (*entity.Story, *apperror.AppError) {
	story, err := u.repo.FindStoryByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Story Not Found")
	}
	if story.UserID != userID {
		return nil, forbidden()
	}
	story.StoryImgUrl = req.StoryImgUrl
	story.Title = req.Title
	story.Description = req.Description
	if err := u.repo.UpdateStory(story); err != nil {
		return nil, apperror.Internal("Failed to update story")
	}
	return story, nil
}

func (u *InvitationUsecase) DeleteStory(userID, id uint) *apperror.AppError {
	story, err := u.repo.FindStoryByID(id)
	if err != nil {
		return mapRepoErr(err, "Story Not Found")
	}
	if story.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteStory(id); err != nil {
		return apperror.Internal("Failed to delete story")
	}
	return nil
}

// ---------- Gift ----------

func (u *InvitationUsecase) CreateGift(userID uint, req entity.GiftRequest) (*entity.Gift, *apperror.AppError) {
	gift := entity.Gift{
		UserID:          userID,
		Provider:        req.Provider,
		ProviderAccount: req.ProviderAccount,
		Type:            req.Type,
		No:              req.No,
	}
	if err := u.repo.CreateGift(&gift); err != nil {
		return nil, apperror.Internal("Failed to create gift")
	}
	return &gift, nil
}

func (u *InvitationUsecase) UpdateGift(userID, id uint, req entity.GiftRequest) (*entity.Gift, *apperror.AppError) {
	gift, err := u.repo.FindGiftByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Gift Not Found")
	}
	if gift.UserID != userID {
		return nil, forbidden()
	}
	gift.Provider = req.Provider
	gift.ProviderAccount = req.ProviderAccount
	gift.Type = req.Type
	gift.No = req.No
	if err := u.repo.UpdateGift(gift); err != nil {
		return nil, apperror.Internal("Failed to update gift")
	}
	return gift, nil
}

func (u *InvitationUsecase) DeleteGift(userID, id uint) *apperror.AppError {
	gift, err := u.repo.FindGiftByID(id)
	if err != nil {
		return mapRepoErr(err, "Gift Not Found")
	}
	if gift.UserID != userID {
		return forbidden()
	}
	if err := u.repo.DeleteGift(id); err != nil {
		return apperror.Internal("Failed to delete gift")
	}
	return nil
}
