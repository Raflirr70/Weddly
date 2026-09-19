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
	ensureSections(&inv.Sections)
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

// ponytail: sectionOrder/sectionEnabled default semua on; keseluruhan tersimpan di kolom JSON Sections.
var sectionOrder = []string{"opening", "invitation", "eventTime", "galery", "storySection", "gift"}

func ensureSections(s *entity.InvitationSections) {
	if s.SectionOrder == nil {
		s.SectionOrder = append([]string{}, sectionOrder...)
	}
	if s.SectionEnabled == nil {
		s.SectionEnabled = make(map[string]bool, len(sectionOrder))
		for _, name := range sectionOrder {
			s.SectionEnabled[name] = true
		}
	}
}

func (u *InvitationUsecase) save(inv *entity.Invitation, invalidMsg string) *apperror.AppError {
	if err := u.repo.UpdateInvitation(inv); err != nil {
		return apperror.Internal(invalidMsg)
	}
	return nil
}

// ---------- Cover ----------

func (u *InvitationUsecase) CreateCover(userID uint, req entity.CoverRequest) (*entity.Cover, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	inv.Sections.CoverUrl = req.CoverUrl
	if appErr := u.save(inv, "Failed to save cover"); appErr != nil {
		return nil, appErr
	}
	return &entity.Cover{CoverUrl: inv.Sections.CoverUrl}, nil
}

func (u *InvitationUsecase) DeleteCover(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	inv.Sections.CoverUrl = ""
	return u.save(inv, "Failed to delete cover")
}

// ---------- Hero ----------

func (u *InvitationUsecase) CreateHero(userID uint, req entity.HeroRequest) (*entity.Hero, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	inv.Sections.HeroImgUrl = req.HeroImgUrl
	if appErr := u.save(inv, "Failed to save hero"); appErr != nil {
		return nil, appErr
	}
	return &entity.Hero{HeroImgUrl: inv.Sections.HeroImgUrl}, nil
}

func (u *InvitationUsecase) DeleteHero(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	inv.Sections.HeroImgUrl = ""
	return u.save(inv, "Failed to delete hero")
}

// ---------- Opening ----------

func (u *InvitationUsecase) CreateOpening(userID uint, req entity.OpeningRequest) (*entity.Opening, *apperror.AppError) {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	inv.Sections.Opening = entity.Opening{
		OpeningImgUrl: req.OpeningImgUrl,
		Title:         req.Title,
		Description:   req.Description,
	}
	if appErr := u.save(inv, "Failed to save opening"); appErr != nil {
		return nil, appErr
	}
	opening := inv.Sections.Opening
	return &opening, nil
}

func (u *InvitationUsecase) DeleteOpening(userID uint) *apperror.AppError {
	inv, appErr := u.invitationByUserID(userID)
	if appErr != nil {
		return appErr
	}
	inv.Sections.Opening = entity.Opening{}
	return u.save(inv, "Failed to delete opening")
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
	inv.Title = req.Title
	events := make([]entity.Event, 0, len(req.Events))
	for _, e := range req.Events {
		events = append(events, entity.Event{
			Order:        e.Order,
			Title:        e.Title,
			Location:     e.Location,
			StartDate:    e.StartDate,
			EndDate:      e.EndDate,
			LocationLink: e.LocationLink,
		})
	}
	inv.Sections.Events = events
	if appErr := u.save(inv, "Failed to save event"); appErr != nil {
		return nil, appErr
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
		galleries = append(galleries, entity.Gallery{Order: g.Order, ImageUrl: g.ImageUrl})
	}
	inv.Sections.Galleries = galleries
	if appErr := u.save(inv, "Failed to save gallery"); appErr != nil {
		return nil, appErr
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
			Order:       s.Order,
			StoryImgUrl: req.StoryImgUrl,
			Title:       s.Title,
			Description: s.Description,
		})
	}
	inv.Sections.Stories = stories
	if appErr := u.save(inv, "Failed to save story"); appErr != nil {
		return nil, appErr
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
			Order:           g.Order,
			Provider:        g.Provider,
			ProviderAccount: g.ProviderAccount,
			Type:            g.Type,
			No:              g.No,
		})
	}
	inv.Sections.Gifts = gifts
	if appErr := u.save(inv, "Failed to save gift"); appErr != nil {
		return nil, appErr
	}
	return gifts, nil
}

// ---------- Public Guest ----------

func (u *InvitationUsecase) GetInvitation(id uint) (*entity.InvitationDetailResponse, *apperror.AppError) {
	inv, err := u.repo.FindInvitationByID(id)
	if err != nil {
		return nil, mapRepoErr(err, "Invitation Not Found")
	}
	ensureSections(&inv.Sections)

	resp := entity.InvitationDetailResponse{
		ID:             inv.ID,
		UserID:         inv.UserID,
		BrideName:      inv.BrideName,
		BrideDegree:    inv.BrideDegree,
		GroomName:      inv.GroomName,
		GroomDegree:    inv.GroomDegree,
		CoverUrl:       inv.Sections.CoverUrl,
		HeroImgUrl:     inv.Sections.HeroImgUrl,
		SectionOrder:   inv.Sections.SectionOrder,
		SectionEnabled: inv.Sections.SectionEnabled,
		Opening: entity.OpeningResponse{
			OpeningImageUrl: inv.Sections.Opening.OpeningImgUrl,
			Title:           inv.Sections.Opening.Title,
			Description:     inv.Sections.Opening.Description,
		},
		Invitation: entity.InvitationSectionResponse{
			GroomImgUrl:      inv.GroomImgUrl,
			BrideImgUrl:      inv.BrideImgUrl,
			Title:            inv.Title,
			Description:      inv.Description,
			GroomDescription: inv.GroomDescription,
			BrideDescription: inv.BrideDescription,
		},
		EventTime: make([]entity.EventResponse, 0, len(inv.Sections.Events)),
		Galery:    make([]entity.GalleryResponse, 0, len(inv.Sections.Galleries)),
		StorySection: entity.StorySectionResponse{
			Story: make([]entity.StoryEntry, 0, len(inv.Sections.Stories)),
		},
		Gift: make([]entity.GiftResponse, 0, len(inv.Sections.Gifts)),
	}

	for _, e := range inv.Sections.Events {
		resp.EventTime = append(resp.EventTime, entity.EventResponse{
			Order:        e.Order,
			Title:        e.Title,
			Location:     e.Location,
			StartDate:    e.StartDate,
			EndDate:      e.EndDate,
			LocationLink: e.LocationLink,
		})
	}
	for _, g := range inv.Sections.Galleries {
		resp.Galery = append(resp.Galery, entity.GalleryResponse{
			Order:    g.Order,
			ImageUrl: g.ImageUrl,
		})
	}
	if len(inv.Sections.Stories) > 0 {
		resp.StorySection.StoryImgUrl = inv.Sections.Stories[0].StoryImgUrl
	}
	for _, s := range inv.Sections.Stories {
		resp.StorySection.Story = append(resp.StorySection.Story, entity.StoryEntry{
			Order:       s.Order,
			Title:       s.Title,
			Description: s.Description,
		})
	}
	for _, g := range inv.Sections.Gifts {
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