package model

import (
	"fmt"
	"gin-blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)))
	if err != nil {
		fmt.Printf("数据库open失败，请检查参数", err)
	}

	//自动迁移您的 schema，保持您的 schema 是最新的
	db.AutoMigrate()

	//转换成*sql.DB类型
	sqlDB, sqlerr := db.DB()
	defer sqlDB.Close()
	if sqlerr != nil {
		fmt.Printf("数据库连接失败", sqlerr)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。一定不能大于框架设置的timeout的时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
