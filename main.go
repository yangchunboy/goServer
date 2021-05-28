package main

import (
	// "gin/router"
	"goServer/router"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()
	router.Router(ginRouter)
	ginRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	ginRouter.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
