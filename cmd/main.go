package main

import (
	"gin--/internal/app"
	"gin--/internal/dao/mysqldb"
)

func main() {
	// 初始化数据库
	mysqldb.InitDB()
	//初始化应用
	app.InitApp()
}
