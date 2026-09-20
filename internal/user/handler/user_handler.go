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

// @Summary      List semua user
// @Tags         User
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /users [get]
func (h *UserHandler) List(c *gin.Context) {
	result, appErr := h.usecase.List()
	if appErr != nil {
		c.JSON(appErr.Code, response.Error(appErr.Code, appErr.Message, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(200, "ok", result))
}

// @Summary      Buat user baru
// @Tags         User
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        body body entity.CreateUserRequest true "Data user"
// @Success      201 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Router       /user [post]
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

// @Summary      Update user (partial)
// @Tags         User
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path int true "User ID"
// @Param        body body entity.UpdateUserRequest true "Field yang diubah"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /user/{id} [put]
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

// @Summary      Hapus user
// @Tags         User
// @Security     BearerAuth
// @Produce      json
// @Param        id   path int true "User ID"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      401 {object} response.Response
// @Failure      403 {object} response.Response
// @Failure      404 {object} response.Response
// @Router       /user/{id} [delete]
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