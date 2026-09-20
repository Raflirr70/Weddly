package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/internal/invitation/entity"
	"github.com/Raflirr70/Weddly/internal/invitation/usecase"
	"github.com/Raflirr70/Weddly/pkg/apperror"
	"github.com/Raflirr70/Weddly/pkg/kafka"
	"github.com/Raflirr70/Weddly/pkg/middleware"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type InvitationHandler struct {
	usecase  *usecase.InvitationUsecase
	producer *kafka.Producer
}

func NewInvitationHandler(usecase *usecase.InvitationUsecase, producer *kafka.Producer) *InvitationHandler {
	return &InvitationHandler{usecase: usecase, producer: producer}
}

func publish(ctx context.Context, p *kafka.Producer, topic, key string, value interface{}, errKey string) {
	if p == nil {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := p.Publish(ctx, topic, key, value); err != nil {
		log.Println(errKey, err)
	}
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

// @Summary      Buat cover undangan
// @Tags         Cover
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.CoverRequest true "URL cover"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /cover [post]
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

// @Summary      Hapus cover undangan (milik user)
// @Tags         Cover
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /cover [delete]
func (h *InvitationHandler) DeleteCover(c *gin.Context) {
	if appErr := h.usecase.DeleteCover(currentUserID(c)); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Hero ----------

// @Summary      Buat hero undangan
// @Tags         Hero
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.HeroRequest true "URL gambar hero"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /hero [post]
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

// @Summary      Hapus hero undangan (milik user)
// @Tags         Hero
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /hero [delete]
func (h *InvitationHandler) DeleteHero(c *gin.Context) {
	if appErr := h.usecase.DeleteHero(currentUserID(c)); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Opening ----------

// @Summary      Buat section opening
// @Tags         Opening
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.OpeningRequest true "Data opening"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /opening [post]
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

// @Summary      Hapus section opening (milik user)
// @Tags         Opening
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /opening [delete]
func (h *InvitationHandler) DeleteOpening(c *gin.Context) {
	if appErr := h.usecase.DeleteOpening(currentUserID(c)); appErr != nil {
		h.fail(c, appErr)
		return
	}
	deleted(c)
}

// ---------- Invitation ----------

// @Summary      Buat undangan
// @Tags         Invitation
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.InvitationRequest true "Data undangan"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /invitation [post]
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

// @Summary      Update undangan
// @Tags         Invitation
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path int true "Invitation ID"
// @Param        body body entity.InvitationRequest true "Data undangan"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /invitation/{id} [put]
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

// @Summary      Hapus undangan
// @Tags         Invitation
// @Security     BearerAuth
// @Produce      json
// @Param        id   path int true "Invitation ID"
// @Success      200 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /invitation/{id} [delete]
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

// @Summary      Buat event (batch)
// @Tags         Event
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.EventRequest true "List event"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /event [post]
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

// ---------- Gallery ----------

// @Summary      Buat gallery (batch)
// @Tags         Gallery
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.GalleryRequest true "List gambar"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /gallery [post]
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

// ---------- Story ----------

// @Summary      Buat story (batch)
// @Tags         Story
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.StoryRequest true "Gambar + list cerita"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /story [post]
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

// ---------- Gift ----------

// @Summary      Buat gift (batch)
// @Tags         Gift
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.GiftRequest true "List rekening/e-wallet"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /gift [post]
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

// ---------- Public Guest ----------

// @Summary      Data undangan lengkap (publik)
// @Tags         Guest
// @Produce      json
// @Param        id   path int true "Invitation ID"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /invitation/{id} [get]
func (h *InvitationHandler) GetInvitation(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	result, appErr := h.usecase.GetInvitation(id)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	go publish(context.Background(), h.producer, kafka.TopicVisitorLog,
		strconv.FormatUint(uint64(result.ID), 10),
		kafka.VisitorLogEvent{
			InvitationID: result.ID,
			UserID:       result.UserID,
			IPAddress:    c.ClientIP(),
			UserAgent:    c.GetHeader("User-Agent"),
			VisitedAt:    time.Now(),
		},
		"publish visitor-log:")
	success(c, result)
}

// @Summary      List komentar undangan (publik)
// @Tags         Guest
// @Produce      json
// @Param        id   path int true "Invitation ID"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /invitation/{id}/comments [get]
func (h *InvitationHandler) GetComments(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	result, appErr := h.usecase.GetComments(id)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	success(c, result)
}

// @Summary      Kirim komentar (publik)
// @Tags         Guest
// @Accept       json
// @Produce      json
// @Param        id   path int true "Invitation ID"
// @Param        body body entity.CommentRequest true "Isi komentar"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /invitation/{id}/comment [post]
func (h *InvitationHandler) CreateComment(c *gin.Context) {
	id, ok := h.paramID(c)
	if !ok {
		return
	}
	var req entity.CommentRequest
	if !h.bind(c, "Comment Bad Request", &req) {
		return
	}
	result, appErr := h.usecase.CreateComment(id, req)
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	go publish(context.Background(), h.producer, kafka.TopicComment,
		strconv.FormatUint(uint64(result.InvitationID), 10),
		kafka.CommentCreatedEvent{
			InvitationID:     result.InvitationID,
			UserID:           result.UserID,
			Username:         result.Username,
			Comment:          result.Comment,
			ConfirmAttendant: result.ConfirmAttendant,
			CreatedAt:        result.CreatedAt,
		},
		"publish comment-created:")
	success(c, result)
}