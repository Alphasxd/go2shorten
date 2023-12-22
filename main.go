package main

import (
	"fmt"

	"github.com/Alphasxd/go2shorten/handler"
	"github.com/Alphasxd/go2shorten/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortURL(c)
	})

	r.GET("/:short_url", func(c *gin.Context) {
		handler.HandleShortURLRedirect(c)
	})

	// 初始化 redis 客户端
	store.InitializeStore()
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
