package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	CreateAt int64  `gorm:"autoCreateTime"`
}

type Todo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"type:varchar(255);not null"`
	Completed bool   `gorm:"type:bool;not null;default:false"`
	CreateAt  int64  `gorm:"autoCreateTime"`
	UpdateAt  int64  `gorm:"autoUpdateTime"`
	UserID    int    `gorm:"type:int;not null"`
	User      User   `json:"-" gorm:"foreignKey:UserID;references:ID"`
}

type TodoCreate struct {
	Title string `json:"title" gorm:"type:varchar(255);not null"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// testStruct
type UserTest struct {
	Name string `json:"name" form:"name" uri:"name" binding:"required"`
	Age  int    `json:"age" form:"age" uri:"age" binding:"gt=18,lt=30"`
	Sex  string `json:"sex" form:"sex" uri:"sex" binding:"oneof=man woman"`
}
