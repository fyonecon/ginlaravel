package bootstrap

import (
	"ginvel.com/app/Http/Middleware"
	"ginvel.com/config"
	"ginvel.com/routes"
	"os"

	//"ginvel.com/routes/Router"
	"github.com/gin-gonic/gin"
	"log"
)

// App 配置并启动http服务
// 项目访问入口
func App(HttpServer *gin.Engine) {
	serverConfig := config.GetServerConfig()

	// Gin服务
	HttpServer = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(Middleware.StatLatency)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(Middleware.AppData)

	// 拦截应用500报错，使之可视化
	HttpServer.Use(Middleware.AppError500)

	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])

	// 配置模版视图
	HttpServer.LoadHTMLGlob(config.GetViewConfig()["View_Pattern"])

	// 注册必要路由，处理默认路由、静态文件路由、404路由等
	Middleware.RouteMust(HttpServer)

	// 注册其他路由，可以自定义
	routes.RouterRegister(HttpServer)
	//Router.Api(HttpServer) // 面向Api
	//Router.Web(HttpServer) // 面向模版输出

	// 访问网址和端口
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// 终端提示
	log.Println(
		"\n App Success! \n Author: fyonecon | Blog: https://ginlaravel.com \n\n " +
			"访问地址示例>>> \n " +
			"1) 访问首页（模版输出）：http://" + host + " \n " +
			"2) 访问接口（JSON输出）：http://" + host + "/api?name=gl&id=2021 \n " +
			"3) 静态文件输出（文件）：http://" + host + "/favicon.ico \n " +
			"")

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误，运行中断：", err.Error())
		os.Exit(200)
	}

	return
}