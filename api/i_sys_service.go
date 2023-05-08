package api

import (
	"go-web/model"

	"github.com/gin-gonic/gin"
)

type ISysService interface {
	GetSysInfo(c *gin.Context) model.SysInfo
}
