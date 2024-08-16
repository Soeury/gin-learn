package main

import (
	"fmt"
	"go_gin/project/dao"
	"go_gin/project/models"
	"go_gin/project/routers"
)

func main() {

	// 创建数据库  create database bubble;
	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Printf("found err in initMysql : %v\n", err)
		return
	}
	defer dao.DB.Close()

	// 建立对应关系
	dao.DB.AutoMigrate(&models.Todo{})
	r := routers.SetupRouter()
	r.Run(":8080")
}
