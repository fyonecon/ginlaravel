package main

import (
	"ginlaravel/bootstrap/app"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"log"
)

var HttpServer *gin.Engine

func main() {
	log.Println("启动服务...")

	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 检测Redis服务
	//driver.GoRedis()

	// 启动服务
	app.Run(HttpServer)
}
