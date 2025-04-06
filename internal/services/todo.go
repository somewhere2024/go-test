package services

import (
	"gin--/internal/dao/mysqldb"
	"gin--/internal/models"
)

func ListUserTodos(id int) {
	user := &models.User{}

	mysqldb.DB.Where("user_id = ?", id).First(user)
	result := mysqldb.DB.Find(&models.Todo{}, "user_id = ?", id)

}
