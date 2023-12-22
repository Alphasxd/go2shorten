package handler

import (
	"net/http"

	"github.com/Alphasxd/go2shorten/shortener"
	"github.com/Alphasxd/go2shorten/store"
	"github.com/gin-gonic/gin"
)

type URLCreationRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest URLCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortURL := shortener.GenerateShortURL(creationRequest.OriginalURL, creationRequest.UserID)
	store.SaveUrlMapping(shortURL, creationRequest.OriginalURL, creationRequest.UserID)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "Short URL created successfully!",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("short_url")
	initialURL := store.RetrieveInitialUrl(shortURL)
	c.Redirect(http.StatusFound, initialURL)
}
