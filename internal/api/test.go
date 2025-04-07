package api

import (
	"fmt"
	"gin--/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHeader(c *gin.Context) {
	header := c.Request.Header
	fmt.Println(header)

	c.JSON(200, gin.H{"message": "test", "data": header})
}

func GetPath(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindUri(user)
	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})
}

func GetQuery(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindQuery(user)

	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})

}

func PostForm(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBind(user)

	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})

}

func PostJson(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})
}
