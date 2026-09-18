package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/internal/invitation/entity"
	"github.com/Raflirr70/Weddly/internal/invitation/usecase"
	"github.com/Raflirr70/Weddly/pkg/apperror"
	"github.com/Raflirr70/Weddly/pkg/middleware"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type InvitationHandler struct {
	usecase *usecase.InvitationUsecase
}

func NewInvitationHandler(usecase *usecase.InvitationUsecase) *InvitationHandler {
	return &InvitationHandler{usecase: usecase}
}

func currentUserID(c *gin.Context) uint {
	v, _ := c.Get(middleware.ContextUserID)
	id, _ := v.(uint)
	return id
}

func (h *InvitationHandler) bind(c *gin.Context, msg string, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, msg, err.Error()))
		return false
	}
	return true
}

func (h *InvitationHandler) fail(c *gin.Context, appErr *apperror.AppError) {
	c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, appErr.Detail))
}

func (h *InvitationHandler) paramID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "Bad Request", err.Error()))
		return 0, false
	}
	return uint(id), true
}

func created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, response.Success(201, "ok", data))
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response.Success(200, "ok", data))
}

func deleted(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(204, "ok", nil))
}

// ---------- Cover ----------

func (h *InvitationHandler) CreateCover(c *gin.Context) {
	var req entity.CoverRequest
	if !h.bind(c, "Cover Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateCover(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) DeleteCover(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteCover(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Hero ----------

func (h *InvitationHandler) CreateHero(c *gin.Context) {
	var req entity.HeroRequest
	if !h.bind(c, "Hero Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateHero(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) DeleteHero(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteHero(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Opening ----------

func (h *InvitationHandler) CreateOpening(c *gin.Context) {
	var req entity.OpeningRequest
	if !h.bind(c, "Opening Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateOpening(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) UpdateOpening(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.OpeningRequest
	if !h.bind(c, "Opening Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.UpdateOpening(currentUserID(c), id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

func (h *InvitationHandler) DeleteOpening(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteOpening(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Invitation ----------

func (h *InvitationHandler) CreateInvitation(c *gin.Context) {
	var req entity.InvitationRequest
	if !h.bind(c, "Invitation Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateInvitation(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) UpdateInvitation(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.InvitationRequest
	if !h.bind(c, "Invitation Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.UpdateInvitation(currentUserID(c), id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

func (h *InvitationHandler) DeleteInvitation(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteInvitation(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Event ----------

func (h *InvitationHandler) CreateEvent(c *gin.Context) {
	var req entity.EventRequest
	if !h.bind(c, "Event Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateEvent(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) UpdateEvent(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.EventRequest
	if !h.bind(c, "Event Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.UpdateEvent(currentUserID(c), id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

func (h *InvitationHandler) DeleteEvent(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteEvent(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Gallery ----------

func (h *InvitationHandler) CreateGallery(c *gin.Context) {
	var req entity.GalleryRequest
	if !h.bind(c, "Gallery Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateGallery(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) DeleteGallery(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteGallery(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Story ----------

func (h *InvitationHandler) CreateStory(c *gin.Context) {
	var req entity.StoryRequest
	if !h.bind(c, "Story Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateStory(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) UpdateStory(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.StoryRequest
	if !h.bind(c, "Story Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.UpdateStory(currentUserID(c), id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

func (h *InvitationHandler) DeleteStory(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteStory(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Gift ----------

func (h *InvitationHandler) CreateGift(c *gin.Context) {
	var req entity.GiftRequest
	if !h.bind(c, "Gift Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateGift(currentUserID(c), req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	created(c, result)
}

func (h *InvitationHandler) UpdateGift(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.GiftRequest
	if !h.bind(c, "Gift Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.UpdateGift(currentUserID(c), id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

func (h *InvitationHandler) DeleteGift(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	if appErr := h.usecase.DeleteGift(currentUserID(c), id); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}
