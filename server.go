package main

import (
	"ginlaravel/app/provider/driver"
	"ginlaravel/app/provider/door"
	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func main() {
	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 启动服务
	door.Run(HttpServer)
}
