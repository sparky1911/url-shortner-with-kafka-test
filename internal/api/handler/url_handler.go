package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your-username/url-shortener/internal/service"
)

type URLHandler struct {
	svc *service.URLService
}

func NewURLHandler(svc *service.URLService) *URLHandler {
	return &URLHandler{svc: svc}
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		LongURL string `json:"url" binding:"required,url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	code, err := h.svc.ShortenUrl(c.Request.Context(), req.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + code})
}

func (h *URLHandler) Redirect(c *gin.Context) {

	code := c.Param("code")
	ip:=c.ClientIP()
	ua:=c.Request.UserAgent()
	longURL, err := h.svc.GetOriginalURL(c.Request.Context(), code,ip,ua)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longURL)

}
