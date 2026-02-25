package api

import (
	"github.com/gin-gonic/gin"
	"github.com/your-username/url-shortener/internal/api/handler"
)

func SetupRouter(h *handler.URLHandler) *gin.Engine {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/shorten", h.Shorten)
	r.GET("/:code", h.Redirect)

	return r
}