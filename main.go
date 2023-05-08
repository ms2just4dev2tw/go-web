package main

import (
	"flag"

	"github.com/gin-gonic/gin"

	"go-web/api"
	"go-web/router"
)

var (
	// 私钥
	secretKey string
	// 指定数据库服务
	dbServer string
	// 数据库名
	dbName string
	// 数据库用户名
	dbUser string
	// 数据库用户密码
	dbPassword string
)

// go 内置函数，先于 main 函数
func init() {
	flag.StringVar(&secretKey, "key", "default-secret-key", "secret key")
	flag.StringVar(&dbServer, "dbserver", "", "database server")
	flag.StringVar(&dbName, "dbname", "", "database name")
	flag.StringVar(&dbUser, "dbuser", "", "database user")
	flag.StringVar(&dbPassword, "dbpwd", "", "database password")
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// init user service
	var userService api.IUserService = InitUserService()
	router.InitUserRouter(r, userService)
	// init sys service
	var sysService api.ISysService = InitSysService()
	router.InitSysRouter(r, sysService)

	return r
}

func main() {
	flag.Parse()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
