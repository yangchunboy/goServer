package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	userId := ctx.Query("userId")
	ctx.String(http.StatusOK, "userId是 %s ", userId)
}

func InsertUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	ctx.JSON(200, gin.H{
		"status":   "posted",
		"username": username,
	})
}
