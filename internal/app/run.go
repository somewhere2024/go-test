package app

import (
	"gin--/config"
	"gin--/internal/routers"
	"gin--/internal/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	R *gin.Engine
)

func InitApp() {
	// 初始化logger
	logger.InitLogger(config.Cfg)
	routers.DefaultRouter()
}
