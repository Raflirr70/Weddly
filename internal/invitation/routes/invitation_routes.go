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
		secured.DELETE("/cover/:id", r.handler.DeleteCover)

		secured.POST("/hero", r.handler.CreateHero)
		secured.DELETE("/hero/:id", r.handler.DeleteHero)

		secured.POST("/opening", r.handler.CreateOpening)
		secured.PUT("/opening/:id", r.handler.UpdateOpening)
		secured.DELETE("/opening/:id", r.handler.DeleteOpening)

		secured.POST("/invitation", r.handler.CreateInvitation)
		secured.PUT("/invitation/:id", r.handler.UpdateInvitation)
		secured.DELETE("/invitation/:id", r.handler.DeleteInvitation)

		secured.POST("/event", r.handler.CreateEvent)
		secured.PUT("/event/:id", r.handler.UpdateEvent)
		secured.DELETE("/event/:id", r.handler.DeleteEvent)

		secured.POST("/gallery", r.handler.CreateGallery)
		secured.DELETE("/gallery/:id", r.handler.DeleteGallery)

		secured.POST("/story", r.handler.CreateStory)
		secured.PUT("/story/:id", r.handler.UpdateStory)
		secured.DELETE("/story/:id", r.handler.DeleteStory)

		secured.POST("/gift", r.handler.CreateGift)
		secured.PUT("/gift/:id", r.handler.UpdateGift)
		secured.DELETE("/gift/:id", r.handler.DeleteGift)
	}

	public := api.Group("")
	{
		public.GET("/invitation/:id", r.handler.GetInvitation)
		public.GET("/invitation/:id/comments", r.handler.GetComments)
		public.POST("/invitation/:id/comment", r.handler.CreateComment)
	}
}
