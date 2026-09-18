package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/Raflirr70/Weddly/internal/notification/handler"
	"github.com/Raflirr70/Weddly/pkg/middleware"
)

type NotificationRoutes struct {
	handler   *handler.NotificationHandler
	redis     *redis.Client
	jwtSecret string
}

func NewNotificationRoutes(h *handler.NotificationHandler, redis *redis.Client, jwtSecret string) *NotificationRoutes {
	return &NotificationRoutes{handler: h, redis: redis, jwtSecret: jwtSecret}
}

func (r *NotificationRoutes) Register(api *gin.RouterGroup) {
	secured := api.Group("")
	secured.Use(
		middleware.AuthMiddleware(r.redis, r.jwtSecret),
		middleware.RoleMiddleware("superadmin"),
	)
	{
		secured.GET("/logs", r.handler.ListLogs)
		secured.GET("/visitors", r.handler.VisitorStats)
	}
}
