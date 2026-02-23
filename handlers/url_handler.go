package handlers

import (
	"net/http"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *services.URLService
}

func NewURLHandler(service *services.URLService) *URLHandler {
	return &URLHandler{service: service}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing URL"})
		return
	}

	code, err := h.service.Shorten(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error shortening URL"})
		return
	}

	host := c.Request.Host
	c.JSON(http.StatusCreated, gin.H{
		"original_url": req.URL,
		"short_url":    "http://" + host + "/" + code,
		"code":         code,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	originalURL, err := h.service.Resolve(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
