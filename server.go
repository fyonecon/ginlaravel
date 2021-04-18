package main

import (
	"fmt"
	"ginlaravel/bootstrap/app"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	fmt.Println("启动服务...")

	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 启动服务
	app.Run(HttpServer)
}
