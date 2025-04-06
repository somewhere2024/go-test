package api

import (
	"gin--/internal/dao/mysqldb"
	"gin--/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

/*type Todo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(255);not null"`
	Completed bool   `gorm:"type:tinyint(1);not null"`
	CreateAt  int64  `gorm:"autoCreateTime"`
	UpdateAt  int64  `gorm:"autoUpdateTime"`
	UserID    int    `gorm:"type:int;not null"`
	User      User   `gorm:"foreignKey:UserID;references:ID"`
}
*/

func TodoCreate(c *gin.Context) {
	user, ok := c.Get("me")
	if !ok || user == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	title := c.PostForm("title")
	claims, ok := user.(*jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	todo := &models.Todo{
		Title:     title,
		Completed: false,
		UserID:    int((*claims)["id"].(float64)),
	}

	result := mysqldb.DB.Create(todo)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "创建成功"})
}

func TodoList(c *gin.Context) {
	me, ok := c.Get("me")
	if !ok || me == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	claims, ok := me.(*jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	user := models.UserInfo{
		Id:       int((*claims)["id"].(float64)),
		Username: (*claims)["username"].(string),
	}
	
}
