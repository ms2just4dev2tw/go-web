package service

import (
	"go-web/model"

	"github.com/gin-gonic/gin"
)

type SysService struct {
}

func (service *SysService) GetSysInfo(c *gin.Context) model.SysInfo {
	var sysInfo = model.SysInfo{}

	return sysInfo
}

// wire Provider
func NewSysService() *SysService {
	return &SysService{}
}
