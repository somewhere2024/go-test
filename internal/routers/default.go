package routers

import (
	"gin--/internal/api"
	"gin--/internal/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

var R *gin.Engine

func DefaultRouter() {
	//R = gin.New()
	//R.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	R = gin.Default()
	R.Use(middlewares.CORS())
	// 默认路由
	R.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})
	{
		v1 := R.Group("/v1")
		v1.GET("/test", api.TestDb)
	}

	//	用户认证授权组

	{
		user := R.Group("/user")
		user.POST("/login", api.Login)
		user.POST("register", api.Register)

	}

	R.Run(":8000")
}
