package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/internal/auth/entity"
	"github.com/Raflirr70/Weddly/internal/auth/usecase"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "User Bad Request", err.Error()))
		return
	}

	result, appErr := h.usecase.Login(req.Username, req.Password)
	if appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}

	c.JSON(http.StatusCreated, response.Success(201, "ok", result))
}