package util

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresqlConfig struct {
	DSN                  string // DSN data source name
	PreferSimpleProtocol bool   // disables implicit prepared statement usage
}

var defaultPostgresqlConfig = postgres.Config{
	DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	PreferSimpleProtocol: true,
}

func GetConn1() *gorm.DB {
	return db
}

func initDataBase1() {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	// grom 使用 database/sql 的内置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
