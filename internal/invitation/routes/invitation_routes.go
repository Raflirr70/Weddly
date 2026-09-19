package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/Raflirr70/Weddly/internal/invitation/handler"
	"github.com/Raflirr70/Weddly/pkg/middleware"
)

type InvitationRoutes struct {
	handler   *handler.InvitationHandler
	redis     *redis.Client
	jwtSecret string
}

func NewInvitationRoutes(h *handler.InvitationHandler, redis *redis.Client, jwtSecret string) *InvitationRoutes {
	return &InvitationRoutes{handler: h, redis: redis, jwtSecret: jwtSecret}
}

func (r *InvitationRoutes) Register(api *gin.RouterGroup) {
	secured := api.Group("")
	secured.Use(
		middleware.AuthMiddleware(r.redis, r.jwtSecret),
		middleware.RoleMiddleware("superadmin", "admin"),
	)
	{
		secured.POST("/cover", r.handler.CreateCover)
		secured.DELETE("/cover", r.handler.DeleteCover)

		secured.POST("/hero", r.handler.CreateHero)
		secured.DELETE("/hero", r.handler.DeleteHero)

		secured.POST("/opening", r.handler.CreateOpening)
		secured.DELETE("/opening", r.handler.DeleteOpening)

		secured.POST("/invitation", r.handler.CreateInvitation)
		secured.PUT("/invitation/:id", r.handler.UpdateInvitation)
		secured.DELETE("/invitation/:id", r.handler.DeleteInvitation)

		secured.POST("/event", r.handler.CreateEvent)
		secured.POST("/gallery", r.handler.CreateGallery)
		secured.POST("/story", r.handler.CreateStory)
		secured.POST("/gift", r.handler.CreateGift)
	}

	public := api.Group("")
	{
		public.GET("/invitation/:id", r.handler.GetInvitation)
		public.GET("/invitation/:id/comments", r.handler.GetComments)
		public.POST("/invitation/:id/comment", r.handler.CreateComment)
	}
}