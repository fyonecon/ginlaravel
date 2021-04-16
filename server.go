package main

import (
	"fmt"
	"ginlaravel/app/provider/door"
	"ginlaravel/app/provider/driver"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	fmt.Println("启动服务...")

	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 启动服务
	door.Run(HttpServer)
}
