package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
	"turing-resume/models"
)

var DB *gorm.DB

func InitDB() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("turing-resume.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// 获取数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection pool")
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(80)           // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间

	// 自动迁移数据库表结构
	db.AutoMigrate(&models.Cx{}, &models.Cy{})

	DB = db
}
