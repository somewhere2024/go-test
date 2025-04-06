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
	Completed bool   `gorm:"type:tinyint(1);not null"`
	CreateAt  int64  `gorm:"autoCreateTime"`
	UpdateAt  int64  `gorm:"autoUpdateTime"`
	UserID    int    `gorm:"type:int;not null"`
	User      User   `gorm:"foreignKey:UserID;references:ID"`
}
