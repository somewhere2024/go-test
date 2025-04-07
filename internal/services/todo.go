package services

import (
	"errors"
	"gin--/internal/dao/mysqldb"
	"gin--/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ListUserTodos(id int) {
	user := &models.User{}

	mysqldb.DB.Where("user_id = ?", id).First(user)
	//result := mysqldb.DB.Find(&models.Todo{}, "user_id = ?", id)

}

func GetTodoList(id int, todos *[]models.Todo) error {
	if id == 0 {
		return errors.New("id不能为空")
	}
	result := mysqldb.DB.Where("user_id", id).Find(todos)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateTitle(c *gin.Context, title string, toUpdateTitle string) error {
	me, _ := c.Get("me")
	user := me.(*jwt.MapClaims)
	TodoUpdate := &models.Todo{}
	result := mysqldb.DB.Where("title = ? and user_id = ?", title, (*user)["id"]).First(TodoUpdate)
	if result.Error != nil {
		return result.Error
	}
	TodoUpdate.Title = toUpdateTitle
	resultErr := mysqldb.DB.Save(TodoUpdate)
	if resultErr.Error != nil {
		return resultErr.Error
	}
	return nil
}

func UpdateCompleted(c *gin.Context, title string, completed string) error {
	me, _ := c.Get("me")

	user := me.(*jwt.MapClaims)

	todoUpdate := &models.Todo{}
	mysqldb.DB.Where("title = ? and user_id = ?", title, (*user)["id"]).First(todoUpdate)
	todoUpdate.Completed = completed == "true" //如果completed为true，则completed为true，否则改为false
	result := mysqldb.DB.Save(todoUpdate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
