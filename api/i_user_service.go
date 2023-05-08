package api

import (
	"go-web/model"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	Login(c *gin.Context) model.User

	Register(c *gin.Context) model.User

	Logoff(c *gin.Context) model.User
}
