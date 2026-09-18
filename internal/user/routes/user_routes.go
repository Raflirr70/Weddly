package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/Raflirr70/Weddly/internal/user/handler"
	"github.com/Raflirr70/Weddly/pkg/middleware"
)

type UserRoutes struct {
	handler   *handler.UserHandler
	redis     *redis.Client
	jwtSecret string
}

func NewUserRoutes(h *handler.UserHandler, redis *redis.Client, jwtSecret string) *UserRoutes {
	return &UserRoutes{handler: h, redis: redis, jwtSecret: jwtSecret}
}

func (r *UserRoutes) Register(api *gin.RouterGroup) {
	secured := api.Group("")
	secured.Use(
		middleware.AuthMiddleware(r.redis, r.jwtSecret),
		middleware.RoleMiddleware("superadmin"),
	)
	{
		secured.GET("/users", r.handler.List)
		secured.POST("/user", r.handler.Create)
		secured.PUT("/user/:id", r.handler.Update)
		secured.DELETE("/user/:id", r.handler.Delete)
	}
}