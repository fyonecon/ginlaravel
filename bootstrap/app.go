package bootstrap
// Author：fyonecon；Blog：https://ginlaravel.com

import (
	"fmt"
	"ginlaravel/app/Middleware"
	"ginlaravel/config"
	"ginlaravel/routes"
	"github.com/gin-gonic/gin"
)

// App 配置并启动http服务
// 项目访问入口
func App(app *gin.Engine) {
	serverConfig := config.GetServerConfig()

	// Gin服务
	app = gin.Default()

	// 捕捉接口运行耗时（必须排第一）
	app.Use(Middleware.StatLatency)

	// 设置全局ctx参数（必须排第二）
	app.Use(Middleware.AppData)

	// 拦截应用500报错，使之可视化
	app.Use(Middleware.AppError500)

	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])

	// 配置模版视图
	app.LoadHTMLGlob(config.GetViewConfig()["View_Pattern"])

	// 注册必要路由，处理默认路由、静态文件路由、404路由等
	Middleware.RouteMust(app)

	// 注册其他路由，可以自定义
	routes.Api(app) // 面向Api
	routes.Web(app) // 面向模版输出

	// 访问网址和端口
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	// 终端提示
	fmt.Println(
		"Success! \n Author:fyonecon | Blog:https://ginlaravel.com \n\n " +
		"访问地址示例>>> \n " +
		"1) 访问首页（模版输出）：http://" + host + " \n " +
		"2) 访问接口（JSON）：http://" + host + "/api?name=gl&id=2021 \n " +
		"3) 静态文件输出：http://" + host + "/favicon.ico \n " +
		"")

	// 启动http服务
	err := app.Run(host)
	if err != nil {
		panic("Run Error: " + err.Error())
	}

}
