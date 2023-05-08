//go:build wireinject
// +build wireinject

package main

import (
	"go-web/repository"
	"go-web/service"

	"github.com/google/wire"
)

func InitUserService() *service.UserService {
	wire.Build(repository.NewUserDao, service.NewUserService)
	return &service.UserService{}
}

func InitSysService() *service.SysService {
	wire.Build(service.NewSysService)
	return &service.SysService{}
}
