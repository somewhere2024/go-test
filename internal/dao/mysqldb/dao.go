package mysqldb

import (
	"gin--/internal/models"
	"gin--/internal/utils/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//连接数据库
//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func InitDB() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	databaseURL := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{}) // &gorm.Config{}参数为可选， 用来给gorm设置一些配置
	if err != nil {
		logger.Logger.Panic("failed to connect database", zap.Error(err))
	}
	mysqldb, err_x := DB.DB()
	if err_x != nil {
		logger.Logger.Panic("failed to connect database", zap.Error(err))
	}
	mysqldb.SetConnMaxLifetime(time.Hour * 1)
	mysqldb.SetMaxIdleConns(10)
	mysqldb.SetMaxOpenConns(100)

	//迁移数据库
	DB.AutoMigrate(&models.User{}, &models.Todo{})
}
