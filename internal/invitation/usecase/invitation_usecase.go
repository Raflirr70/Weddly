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

func (u *InvitationUsecase) invitationByUserID(userID uint) (*entity.Invitation, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByUserID(userID)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}
	return inv, nil
}

func (u *InvitationUsecase) checkOwner(userID, invitationID uint) *apperror.AppError {
	inv, err := u.repo.FindInvitationByID(invitationID)
	if err != nil {
		return mapRepoErr(err, "Invitation Not Found")
	}
	if inv.UserID != userID {
		return forbidden()
	}
	return nil
}

// sectionOrder: urutan tampil default; tidak dipersist (diatur klien).
var sectionOrder = []string{"opening", "invitation", "eventTime", "galery", "storySection", "gift"}

// ---------- Cover ----------

func (u *InvitationUsecase) CreateCover(userID uint, req entity.CoverRequest) (*entity.Cover, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	cover := &entity.Cover{CoverUrl: req.CoverUrl}
	if err := u.repo.UpsertCover(inv.ID, cover); err != nil {
		return nil, apperror.Internal("Failed to save cover")
	}
	return cover, nil
}

func (u *InvitationUsecase) DeleteCover(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	if err := u.repo.DeleteCover(inv.ID); err != nil {
		return apperror.Internal("Failed to delete cover")
	}
	return nil
}

// ---------- Hero ----------

func (u *InvitationUsecase) CreateHero(userID uint, req entity.HeroRequest) (*entity.Hero, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	hero := &entity.Hero{HeroImgUrl: req.HeroImgUrl}
	if err := u.repo.UpsertHero(inv.ID, hero); err != nil {
		return nil, apperror.Internal("Failed to save hero")
	}
	return hero, nil
}

func (u *InvitationUsecase) DeleteHero(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	if err := u.repo.DeleteHero(inv.ID); err != nil {
		return apperror.Internal("Failed to delete hero")
	}
	return nil
}

// ---------- Opening ----------

func (u *InvitationUsecase) CreateOpening(userID uint, req entity.OpeningRequest) (*entity.Opening, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	opening := &entity.Opening{
		OpeningImgUrl: req.OpeningImgUrl,
		Title:         req.Title,
		Description:   req.Description,
	}
	if err := u.repo.UpsertOpening(inv.ID, opening); err != nil {
		return nil, apperror.Internal("Failed to save opening")
	}
	return opening, nil
}

func (u *InvitationUsecase) DeleteOpening(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	if err := u.repo.DeleteOpening(inv.ID); err != nil {
		return apperror.Internal("Failed to delete opening")
	}
	return nil
}

// ---------- Invitation ----------

func (u *InvitationUsecase) CreateInvitation(userID uint, req entity.InvitationRequest) (*entity.Invitation, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByUserID(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperror.Internal("Failed to get invitation")
	}
	create := err != nil
	if create {
		inv = &entity.Invitation{UserID: userID}
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
	if create {
		if err := u.repo.CreateInvitation(inv); err != nil {
			return nil, apperror.Internal("Failed to create invitation")
		}
	} else if err := u.repo.UpdateInvitation(inv); err != nil {
		return nil, apperror.Internal("Failed to update invitation")
	}
	return inv, nil
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

func (u *InvitationUsecase) CreateEvent(userID uint, req entity.EventRequest) ([]entity.Event, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	events := make([]entity.Event, 0, len(req.Events))
	for _, e := range req.Events {
		events = append(events, entity.Event{
			InvitationID: inv.ID,
			Order:        e.Order,
			Title:        e.Title,
			Location:     e.Location,
			StartDate:    e.StartDate,
			EndDate:      e.EndDate,
			LocationLink: e.LocationLink,
		})
	}
	if err := u.repo.ReplaceEvents(inv.ID, events); err != nil {
		return nil, apperror.Internal("Failed to save event")
	}
	return events, nil
}

// ---------- Gallery ----------

func (u *InvitationUsecase) CreateGallery(userID uint, req entity.GalleryRequest) ([]entity.Gallery, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	galleries := make([]entity.Gallery, 0, len(req.Galleries))
	for _, g := range req.Galleries {
		galleries = append(galleries, entity.Gallery{
			InvitationID: inv.ID,
			Order:        g.Order,
			ImageUrl:     g.ImageUrl,
		})
	}
	if err := u.repo.ReplaceGalleries(inv.ID, galleries); err != nil {
		return nil, apperror.Internal("Failed to save gallery")
	}
	return galleries, nil
}

// ---------- Story ----------

func (u *InvitationUsecase) CreateStory(userID uint, req entity.StoryRequest) ([]entity.Story, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	stories := make([]entity.Story, 0, len(req.Stories))
	for _, s := range req.Stories {
		stories = append(stories, entity.Story{
			InvitationID: inv.ID,
			Order:        s.Order,
			StoryImgUrl:  req.StoryImgUrl,
			Title:        s.Title,
			Description:  s.Description,
		})
	}
	if err := u.repo.ReplaceStories(inv.ID, stories); err != nil {
		return nil, apperror.Internal("Failed to save story")
	}
	return stories, nil
}

// ---------- Gift ----------

func (u *InvitationUsecase) CreateGift(userID uint, req entity.GiftRequest) ([]entity.Gift, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	gifts := make([]entity.Gift, 0, len(req.Gifts))
	for _, g := range req.Gifts {
		gifts = append(gifts, entity.Gift{
			InvitationID:    inv.ID,
			Order:           g.Order,
			Provider:        g.Provider,
			ProviderAccount: g.ProviderAccount,
			Type:            g.Type,
			No:              g.No,
		})
	}
	if err := u.repo.ReplaceGifts(inv.ID, gifts); err != nil {
		return nil, apperror.Internal("Failed to save gift")
	}
	return gifts, nil
}

// ---------- Public Guest ----------

func (u *InvitationUsecase) GetInvitation(id uint) (*entity.InvitationDetailResponse, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}

	events, err := u.repo.GetEvents(inv.ID)
	if err != nil {
		return nil, apperror.Internal("Failed to load events")
	}
	galleries, err := u.repo.GetGalleries(inv.ID)
	if err != nil {
		return nil, apperror.Internal("Failed to load galleries")
	}
	stories, err := u.repo.GetStories(inv.ID)
	if err != nil {
		return nil, apperror.Internal("Failed to load stories")
	}
	gifts, err := u.repo.GetGifts(inv.ID)
	if err != nil {
		return nil, apperror.Internal("Failed to load gifts")
	}

	resp := entity.InvitationDetailResponse{
		ID:          inv.ID,
		UserID:      inv.UserID,
		BrideName:   inv.BrideName,
		BrideDegree: inv.BrideDegree,
		GroomName:   inv.GroomName,
		GroomDegree: inv.GroomDegree,
		SectionOrder:   append([]string{}, sectionOrder...),
		SectionEnabled: make(map[string]bool, len(sectionOrder)),
		Invitation: entity.InvitationSectionResponse{
			GroomImgUrl:      inv.GroomImgUrl,
			BrideImgUrl:      inv.BrideImgUrl,
			Title:            inv.Title,
			Description:      inv.Description,
			GroomDescription: inv.GroomDescription,
			BrideDescription: inv.BrideDescription,
		},
		EventTime:     make([]entity.EventResponse, 0, len(events)),
		Galery:        make([]entity.GalleryResponse, 0, len(galleries)),
		StorySection:  entity.StorySectionResponse{Story: make([]entity.StoryEntry, 0, len(stories))},
		Gift:          make([]entity.GiftResponse, 0, len(gifts)),
	}
	for _, name := range sectionOrder {
		resp.SectionEnabled[name] = true
	}

	if cover, err := u.repo.GetCover(inv.ID); err == nil {
		resp.CoverUrl = cover.CoverUrl
	}
	if hero, err := u.repo.GetHero(inv.ID); err == nil {
		resp.HeroImgUrl = hero.HeroImgUrl
	}
	if opening, err := u.repo.GetOpening(inv.ID); err == nil {
		resp.Opening = entity.OpeningResponse{
			OpeningImageUrl: opening.OpeningImgUrl,
			Title:           opening.Title,
			Description:     opening.Description,
		}
	}

	for _, e := range events {
		resp.EventTime = append(resp.EventTime, entity.EventResponse{
			Order:        e.Order,
			Title:        e.Title,
			Location:     e.Location,
			StartDate:    e.StartDate,
			EndDate:      e.EndDate,
			LocationLink: e.LocationLink,
		})
	}
	for _, g := range galleries {
		resp.Galery = append(resp.Galery, entity.GalleryResponse{
			Order:    g.Order,
			ImageUrl: g.ImageUrl,
		})
	}
	if len(stories) > 0 {
		resp.StorySection.StoryImgUrl = stories[0].StoryImgUrl
	}
	for _, s := range stories {
		resp.StorySection.Story = append(resp.StorySection.Story, entity.StoryEntry{
			Order:       s.Order,
			Title:       s.Title,
			Description: s.Description,
		})
	}
	for _, g := range gifts {
		resp.Gift = append(resp.Gift, entity.GiftResponse{
			Order:           g.Order,
			Provider:        g.Provider,
			ProviderAccount: g.ProviderAccount,
			Type:            g.Type,
			No:              g.No,
		})
	}
	return &resp, nil
}

func (u *InvitationUsecase) GetComments(id uint) ([]entity.CommentResponse, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}
	comments, err := u.repo.FindCommentsByUserID(inv.UserID)
	if err != nil {
		return nil, apperror.Internal("Failed to list comments")
	}
	result := make([]entity.CommentResponse, 0, len(comments))
	for _, c := range comments {
		result = append(result, entity.CommentResponse{
			Username:         c.Username,
			Comment:          c.Comment,
			ConfirmAttendant: c.ConfirmAttendant,
			CreatedAt:        c.CreatedAt,
		})
	}
	return result, nil
}

func (u *InvitationUsecase) CreateComment(id uint, req entity.CommentRequest) (*entity.CommentResponse, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}
	comment := entity.Comment{
		UserID:           inv.UserID,
		Username:         req.Username,
		Comment:          req.Comment,
		ConfirmAttendant: req.ConfirmAttendant,
	}
	if err := u.repo.CreateComment(&comment); err != nil {
		return nil, apperror.Internal("Failed to create comment")
	}
	return &entity.CommentResponse{
		InvitationID:     inv.ID,
		UserID:           inv.UserID,
		Username:         comment.Username,
		Comment:          comment.Comment,
		ConfirmAttendant: comment.ConfirmAttendant,
		CreatedAt:        comment.CreatedAt,
	}, nil
}