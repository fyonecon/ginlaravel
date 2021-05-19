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
func App(httpServer *gin.Engine) {
	serverConfig := config.GetServerConfig()

	// Gin服务
	httpServer = gin.Default()
	// 拦截服务器500报错，使之可视化
	httpServer.Use(Middleware.Server500)
	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])
	// 配置模版视图
	httpServer.LoadHTMLGlob(config.GetViewConfig()["View_Pattern"])
	// 注册路由
	routes.Must(httpServer) // 必要路由
	routes.Web(httpServer) // 面向模版输出
	routes.Api(httpServer) // 面向Api
	// 访问网址和端口
	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]
	// 终端提示
	fmt.Println(
		" Author：fyonecon \n Blog；https://ginlaravel.com \n\n " +
		"访问地址示例>>> \n " +
		"0) 访问首页：http://" + serverAddr + " \n " +
		"1) tpl模版视图输出：http://" + serverAddr + "/tpl?name=danchaofan&id=1949 \n " +
		"2) api前后端分离：http://" + serverAddr + "/api?name=gl&id=2021 \n " +
		"3) 静态文件输出：http://" + serverAddr + "/favicon.ico \n " +
		"4) Swagger接口文档：http://" + serverAddr + "/swagger/index.html \n " +
		"")
	// 启动http服务
	err := httpServer.Run(serverAddr)
	if err != nil {
		panic("Run Error: " + err.Error())
	}

}
