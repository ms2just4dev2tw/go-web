package router

import (
	"go-web/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.Engine, service api.IUserService) {
	// user router group
	userGroup := router.Group("/user")
	{
		userGroup.POST("login", func(ctx *gin.Context) {
			user := service.Login(ctx)

			//
			ctx.JSON(http.StatusOK, gin.H{"user": user})
		})

		userGroup.POST("register", func(ctx *gin.Context) {
			user := service.Register(ctx)

			//
			ctx.JSON(http.StatusOK, gin.H{"user": user})
		})

		userGroup.POST("logoff", func(ctx *gin.Context) {
			user := service.Logoff(ctx)

			//
			ctx.JSON(http.StatusOK, gin.H{"user": user})
		})
	}
}
