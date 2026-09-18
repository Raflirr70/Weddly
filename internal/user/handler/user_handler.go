package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/internal/user/entity"
	"github.com/Raflirr70/Weddly/internal/user/usecase"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) List(c *gin.Context) {
	result, appErr := h.usecase.List()
	if appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(200, "ok", result))
}

func (h *UserHandler) Create(c *gin.Context) {
	var req entity.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "User Bad Request", err.Error()))
		return
	}

	result, appErr := h.usecase.Create(req)
	if appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}
	c.JSON(http.StatusCreated, response.Success(201, "ok", result))
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "User Bad Request", err.Error()))
		return
	}

	var req entity.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "User Bad Request", err.Error()))
		return
	}

	result, appErr := h.usecase.Update(uint(id), req)
	if appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(200, "ok", result))
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(400, "User Bad Request", err.Error()))
		return
	}

	if appErr := h.usecase.Delete(uint(id)); appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}
	// ponytail: DELETE kirim body (code 204) biar sesuai kontrak, bukan 204 no-content.
	c.JSON(http.StatusOK, response.Success(204, "ok", nil))
}