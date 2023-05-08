package router

import (
	"go-web/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitSysRouter(router *gin.Engine, service api.ISysService) {
	// user router group
	sysGroup := router.Group("/sys")
	{
		sysGroup.GET("info", func(ctx *gin.Context) {
			sysInfo := service.GetSysInfo(ctx)

			//
			ctx.JSON(http.StatusOK, gin.H{"sys": sysInfo})
		})
	}
}
