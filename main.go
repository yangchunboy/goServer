package main

import (
	// "gin/router"
	"goServer/router"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()
	router.Router(ginRouter)
	ginRouter.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
