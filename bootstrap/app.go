package bootstrap

import (
	"ginvel.com/app/Http/Middlewares"
	"ginvel.com/app/Ruler/Task"
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
	frameworkConfig := config.GetFrameworkConfig()

	// Gin服务
	HttpServer = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	HttpServer.Use(Middlewares.StatLatency)

	// 设置全局ctx参数（必须排第二）
	HttpServer.Use(Middlewares.AppData)

	// 拦截应用500报错，使之可视化
	HttpServer.Use(Middlewares.AppError500)

	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])

	// 注册必要路由，处理默认路由、静态文件路由、404路由等
	routes.RouteMust(HttpServer)

	// 注册其他路由，可以自定义
	routes.RouterRegister(HttpServer)
	//Router.Api(HttpServer) // 面向Api
	//Router.Web(HttpServer) // 面向模版输出

	// 初始化定时器（立即运行定时器）
	Task.TimeInterval(0, 0, "0")

	// 实际访问网址和端口
	_host := "127.0.0.1:" + serverConfig["PORT"] // 测试访问IP
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"] // Docker访问IP

	glVersion := frameworkConfig["gl_version"]

	// 终端提示
	log.Println(
		//"\n App Success! \n Author: fyonecon | Blog: https://ginlaravel.com \n\n " +
		" \n " +
			"访问地址示例：" + host + " >>> \n " +
			"gl_version：" + glVersion + " \n " +
			"1) 默认接口（JSON输出）：http://" + _host + " \n " +
			"2) 测试接口（JSON输出）：http://" + _host + "/api?name=gl&id=2021 \n " +
			"3) 静态文件输出（文件）：http://" + _host + "/favicon.ico \n " +
			"4) 查看WebSocket连接：ws://" + _host + "/api/example/socket/ping1 \n " +
			"")

	err := HttpServer.Run(host)
	if err != nil {
		log.Println("http服务遇到错误，运行中断，error：", err.Error())
		log.Println("提示：注意端口被占时应该首先更改对外暴露的端口，而不是微服务的端口。")
		os.Exit(200)
	}

	return
}