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
		user := R.Group("/auth")
		user.POST("/login", api.Login)
		user.POST("/register", api.Register)

	}

	//api保护路由
	{
		m := R.Group("/me")
		m.Use(middlewares.Auth())

		m.GET("/readme", api.ReadMe)

	}

	//todo的api
	{
		todo := R.Group("/todo")
		todo.Use(middlewares.Auth()) //保护路由
		todo.POST("/create", api.TodoCreate)
		todo.GET("/test", api.TodoTest)
		todo.GET("/list", api.TodoList)
		todo.GET("/detail/:title", api.TodoDetail)
		todo.DELETE("/delete/:title", api.TodoDelete)
		todo.PUT("/update", api.TodoUpdate)
	}

	//test
	{
		test := R.Group("/test")
		test.GET("/getHeader", api.GetHeader)
		test.POST("/postjson", api.PostJson)
		test.POST("/postform", api.PostForm)
		test.GET("/getquery", api.GetQuery)
		test.GET("/getpath/:name/:age/:sex", api.GetPath)
	}

	R.Run(":8000")
}
