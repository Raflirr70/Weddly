package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"

	"github.com/Raflirr70/Weddly/internal/auth/entity"
	"github.com/Raflirr70/Weddly/pkg/response"
)

const (
	ContextUserID = "userID"
	ContextRole   = "role"

	// ponytail: format key dobel dengan auth_repository.SaveToken.
	tokenKeyPrefix = "auth:token:"
)

func AuthMiddleware(redisClient *redis.Client, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		token, ok := strings.CutPrefix(raw, "Bearer ")
		if !ok {
			token = raw
		}
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(401, "Unauthorized", nil))
			return
		}

		claims := &entity.Claims{}
		parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !parsed.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(401, "Unauthorized", nil))
			return
		}

		key := tokenKeyPrefix + strconv.FormatUint(uint64(claims.UserID), 10)
		stored, err := redisClient.Get(context.Background(), key).Result()
		if err != nil || stored != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error(401, "Unauthorized", nil))
			return
		}

		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextRole, claims.Role)
		c.Next()
	}
}