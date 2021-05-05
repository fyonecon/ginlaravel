package app

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"ginlaravel/config"
	"ginlaravel/routes"
	"github.com/gin-gonic/gin"
	"log"
)

// Run 配置并启动http服务
// 项目访问入口
func Run(httpServer *gin.Engine) {
	// Gin服务
	httpServer = gin.Default()
	// 参数
	serverConfig := config.GetServerConfig()
	// Gin运行时：release、debug、test
	gin.SetMode(serverConfig["ENV"])
	// 配置模版视图
	if serverConfig["VIEWS_PATTERN"] != "" {
		httpServer.LoadHTMLGlob(serverConfig["VIEWS_PATTERN"])
	}
	// 注册路由
	routes.Must(httpServer) // 必要路由
	routes.Api(httpServer) // 面向Api
	routes.Web(httpServer) // 面向模版输出
	// 访问网址和端口
	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]
	// 终端提示
	log.Println("访问地址示例>>> \n tpl模版输出：http://" + serverAddr + "/web/gen1/app/tpl \n api前后端分离：http://" + serverAddr + "/api/gen1/app/api \n")
	// 启动http服务
	err := httpServer.Run(serverAddr)
	if err != nil {
		panic("Run Error: " + err.Error())
	}

}
