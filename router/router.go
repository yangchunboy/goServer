package router

import "github.com/gin-gonic/gin"

func Router(ginRouter *gin.Engine) {

	// 建立接口分组
	v1 := ginRouter.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
