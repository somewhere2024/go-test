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

func TodoTest(c *gin.Context) {
	me, ok := c.Get("me")
	if !ok || me == nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	user, ok := me.(*jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	id := (*user)["id"].(float64)
	username := (*user)["username"].(string)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": gin.H{"id": id, "username": username}})

}

func TodoList(c *gin.Context) {
	me, ok := c.Get("me")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	user, _ := me.(*jwt.MapClaims)
	id := (*user)["id"].(float64)
	todos := &[]models.Todo{}
	err := services.GetTodoList(int(id), todos)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": todos})
}

func TodoUpdate(c *gin.Context) {
	title := c.DefaultPostForm("title", "")
	if title == "" {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "title 不能为空"})
		return
	}
	toUpdateTitle := c.DefaultPostForm("newTitle", "")
	toCompleted := c.DefaultPostForm("completed", "false")
	logger.Logger.Info(fmt.Sprintf("title:%s,newTitle:%s,completed:%s", title, toUpdateTitle, toCompleted))
	if toUpdateTitle == "" && toCompleted != "" {
		result := services.UpdateCompleted(c, title, toCompleted)
		if result != nil {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "更新失败"})
			return
		}
	} else if toCompleted == "false" && toUpdateTitle != "" {
		result := services.UpdateTitle(c, title, toUpdateTitle)
		if result != nil {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "更新失败"})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "title 不能为空 和 completed 不能同时为空"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "update success"})
}

func TodoDetail(c *gin.Context) {
	me, ok := c.Get("me")

	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	todoTitle := c.Param("title") //路径参数获取需要查询的title
	if todoTitle == "" {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "id 不能为空"})
		return
	}
	user := me.(*jwt.MapClaims)
	userId := (*user)["id"].(float64)
	todo, err := services.GetTodo(int(userId), todoTitle)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "查询失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success", "data": todo})
}

func TodoDelete(c *gin.Context) {
	me, ok := c.Get("me")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "message": "token 无效 未授权"})
		return
	}
	todoTitle := c.Param("title")
	if todoTitle == "" {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "message": "id 不能为空"})
		return
	}
	user := me.(*jwt.MapClaims)
	userId := (*user)["id"].(float64)
	err := services.DeleteTodo(int(userId), todoTitle)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "message": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "delete success"})
}
