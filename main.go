package main

import (
	"ginvel.com/bootstrap"
	"ginvel.com/bootstrap/driver"
	"ginvel.com/config"
	"github.com/gin-gonic/gin"
	"log"
)

var HttpServer *gin.Engine

// main
// @title GinLaravel、GL、Ginvel
// @description GinLaravel Web Framework. 2021. In Changsha. Powered By Gin & Laravel & Others, Think You.
// @author https://github.com/fyonecon
// @licence MIT
// @website https://ginlaravel.com
func main() {
	log.Println("【GinLaravel main.go】启动main.go...")
	log.Println("【GOPATH校准 framework.go】当前go_path=", config.GetFrameworkConfig()["go_path"], " \n ")

	// 服务停止时清理数据库链接
	defer driver.MysqlDb.Close()

	// 启动服务
	bootstrap.App(HttpServer)
}
