package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/Raflirr70/Weddly/internal/auth/handler"
	"github.com/Raflirr70/Weddly/pkg/middleware"
	"github.com/Raflirr70/Weddly/pkg/response"
)

type AuthRoutes struct {
	handler   *handler.AuthHandler
	redis     *redis.Client
	jwtSecret string
}

func NewAuthRoutes(h *handler.AuthHandler, redis *redis.Client, jwtSecret string) *AuthRoutes {
	return &AuthRoutes{handler: h, redis: redis, jwtSecret: jwtSecret}
}

func (r *AuthRoutes) Register(api *gin.RouterGroup) {
	api.POST("/login", r.handler.Login)

	secured := api.Group("")
	secured.Use(middleware.AuthMiddleware(r.redis, r.jwtSecret))
	{
		// ponytail: route test middleware, diganti route user CRUD di Step 6.
		secured.GET("/me", me)
		secured.GET("/superadmin", middleware.RoleMiddleware("superadmin"), superadminOnly)
	}
}

func me(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(200, "ok", gin.H{
		"userId": c.GetUint(middleware.ContextUserID),
		"role":   c.GetString(middleware.ContextRole),
	}))
}

func superadminOnly(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success(200, "ok", gin.H{"ok": true}))
}