package main

import (
	"fmt"

	"github.com/Alphasxd/go2shorten/handler"
	"github.com/Alphasxd/go2shorten/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// / 为根路径
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// /createShortUrl 为创建短链接的路径
	r.POST("/createShortUrl", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// /:shortUrl 为短链接重定向的路径
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// 初始化存储
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
