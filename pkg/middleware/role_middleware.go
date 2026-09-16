package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/pkg/response"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(ContextRole)
		roleStr, ok := role.(string)
		if !exists || !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, response.Error(403, "Forbidden", nil))
			return
		}

		for _, r := range roles {
			if roleStr == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, response.Error(403, "Forbidden", nil))
	}
}