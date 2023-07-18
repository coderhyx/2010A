package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL(dsn string) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	// 设置连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取数据库连接失败, error=" + err.Error())
	}
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetMaxIdleConns(10)  // 空闲连接数
}

func GetDB() *gorm.DB {
	return db
}
