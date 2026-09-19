package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Raflirr70/Weddly/pkg/response"
)

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func proxy(targetHost string) gin.HandlerFunc {
	target, _ := url.Parse("http://" + targetHost)
	rp := httputil.NewSingleHostReverseProxy(target)
	return func(c *gin.Context) {
		rp.ServeHTTP(c.Writer, c.Request)
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	authHost := envOr("AUTH_SERVICE_ADDR", "localhost:8081")
	invHost := envOr("INVITATION_SERVICE_ADDR", "localhost:8082")

	authProxy := proxy(authHost)
	invProxy := proxy(invHost)

	routes := []struct {
		prefix string
		next   gin.HandlerFunc
	}{
		{"/api/v1/login", authProxy},
		{"/api/v1/me", authProxy},
		{"/api/v1/user", authProxy},
		{"/api/v1/users", authProxy},
		{"/api/v1/logs", authProxy},
		{"/api/v1/visitors", authProxy},
		{"/api/v1/cover", invProxy},
		{"/api/v1/hero", invProxy},
		{"/api/v1/opening", invProxy},
		{"/api/v1/invitation", invProxy},
		{"/api/v1/event", invProxy},
		{"/api/v1/gallery", invProxy},
		{"/api/v1/story", invProxy},
		{"/api/v1/gift", invProxy},
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), cors(), func(c *gin.Context) {
		path := c.Request.URL.Path
		for _, route := range routes {
			if strings.HasPrefix(path, route.prefix) {
				route.next(c)
				return
			}
		}
		c.JSON(http.StatusNotFound, response.Error(404, "Not Found", nil))
	})

	log.Printf("api-gateway running on :8080 (auth=%s, invitation=%s)", authHost, invHost)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
