package bootstrap

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"fmt"
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
	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])
	// 配置模版视图
	httpServer.LoadHTMLGlob(config.GetViewConfig()["View_Pattern"])
	// 注册路由
	routes.Must(httpServer) // 必要路由
	routes.Api(httpServer) // 面向Api
	routes.Web(httpServer) // 面向模版输出
	// 访问网址和端口
	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]
	// 终端提示
	fmt.Println("\n 访问地址示例>>> \n " +
		"0) 访问首页：http://" + serverAddr + " \n " +
		"1) tpl模版视图输出：http://" + serverAddr + "/tpl?name=danchaofan&id=1949 \n " +
		"2) api前后端分离：http://" + serverAddr + "/api?name=gl&id=2021 \n " +
		"3) 静态文件输出：http://" + serverAddr + "/favicon.ico \n " +
		"")
	// 启动http服务
	err := httpServer.Run(serverAddr)
	if err != nil {
		panic("Run Error: " + err.Error())
	}

}
