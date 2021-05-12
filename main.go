package main
// Author：fyonecon；Blog：https://ginlaravel.com

import (
	"ginlaravel/app/Kit"
	"ginlaravel/bootstrap"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"log"
)

var HttpServer *gin.Engine

func main() {
	log.Println("启动服务...")
	Kit.Log("启动服务", "")

	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 启动服务
	bootstrap.App(HttpServer)
}
