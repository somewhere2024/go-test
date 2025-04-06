package api

import (
	"fmt"
	"gin--/internal/dao/mysqldb"
	"gin--/internal/models"
	"gin--/internal/services"
	"gin--/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func TestDb(c *gin.Context) {
	user := &models.User{
		Username: "test",
		Password: "test",
	}
	db := mysqldb.DB.Create(user)
	fmt.Print(db)
	c.JSON(200, gin.H{"message": "test"})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result, err := services.VerifyUsernamePassword(username, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "用户名或密码错误"})
		return
	}
	if result == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "用户名或密码错误"})
		return
	}
	payload := jwt.MapClaims{
		"id":       result.ID,
		"username": result.Username,
	}
	token, err := services.CreateToken(payload)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "登录成功", "token": token, "token_type": "Bearer"})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		logger.Logger.Warn("用户名或密码不能为空")
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "用户名或密码不能为空"})
		return
	}

	result, err := services.CreateUser(username, password)

	if err != nil {
		logger.Logger.Warn("用户名或密码错误")
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "用户名或密码错误"})
		return
	}
	if result == nil {
		logger.Logger.Warn("用户名或密码错误")
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "用户名或密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "注册成功", "data": gin.H{"username": result.Username, "password": result.Password}})
}
