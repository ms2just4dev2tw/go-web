package service

import (
	"go-web/model"
	"go-web/repository"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	dao *repository.UserDao
}

func (service *UserService) Login(c *gin.Context) model.User {
	var user = model.User{}

	return user
}

func (service *UserService) Register(c *gin.Context) model.User {
	var user = model.User{}

	return user
}

func (service *UserService) Logoff(c *gin.Context) model.User {
	var user = model.User{}

	return user
}

// wire Provider
func NewUserService(dao *repository.UserDao) *UserService {
	return &UserService{dao: dao}
}
