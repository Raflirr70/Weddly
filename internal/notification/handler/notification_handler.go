package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/internal/notification/usecase"
	"github.com/Raflirr70/Weddly/pkg/apperror"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type NotificationHandler struct {
	usecase *usecase.NotificationUsecase
}

func NewNotificationHandler(usecase *usecase.NotificationUsecase) *NotificationHandler {
	return &NotificationHandler{usecase: usecase}
}

func (h *NotificationHandler) fail(c *gin.Context, appErr *apperror.AppError) {
	c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, appErr.Detail))
}

func (h *NotificationHandler) ListLogs(c *gin.Context) {
	result, appErr := h.usecase.ListLogs()
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	c.JSON(http.StatusOK, response.Success(200, "ok", result))
}

func (h *NotificationHandler) VisitorStats(c *gin.Context) {
	result, appErr := h.usecase.VisitorStats()
	if appErr != nil {
		h.fail(c, appErr)
		return
	}
	c.JSON(http.StatusOK, response.Success(200, "ok", result))
}
