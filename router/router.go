package router

import (
	"github.com/gin-gonic/gin"

	"goServer/controller/user"
)

func Router(ginRouter *gin.Engine) {

	// 建立接口分组
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("/user/detail", user.GetUserInfo)
		v1.POST("/user/insert", user.InsertUser)
	}
}
