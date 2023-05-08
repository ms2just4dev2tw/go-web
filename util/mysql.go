package util

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	DSN                       string // DSN data source name
	DefaultStringSize         uint   // string 类型字段的默认长度
	DisableDatetimePrecision  bool   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	DontSupportRenameIndex    bool   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	DontSupportRenameColumn   bool   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	SkipInitializeWithVersion bool   // 根据当前 MySQL 版本自动配置
}

var defaultMysqlConfig = mysql.Config{
	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local",
	DefaultStringSize:         256,
	DisableDatetimePrecision:  true,
	DontSupportRenameIndex:    true,
	DontSupportRenameColumn:   true,
	SkipInitializeWithVersion: false,
}

var db *gorm.DB

func GetConn() *gorm.DB {
	bcrypt.CompareHashAndPassword([]byte("password"), []byte("123"))
	return db
}

func initDataBase() {
	db, _ := gorm.Open(mysql.New(defaultMysqlConfig), &gorm.Config{})
	// grom 使用 database/sql 的内置连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}
