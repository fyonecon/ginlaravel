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

// 配置并启动http服务
// 项目访问入口
func Run(httpServer *gin.Engine) {
	log.Println("检查中间件")

	// 服务配置
	serverConfig := config.GetServerConfig()

	// gin 运行时 release debug test
	gin.SetMode(serverConfig["ENV"])

	httpServer = gin.Default()

	// 配置视图
	if "" != serverConfig["VIEWS_PATTERN"] {
		httpServer.LoadHTMLGlob(serverConfig["VIEWS_PATTERN"])
	}

	// 注册路由
	routes.RegisterRoutes(httpServer)

	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	log.Println("提示：Gin服务内存常驻，请提前使用screen会话服务(need yum install screen)来继续保留终端窗口；退出服务请按：'Ctrl + C' 。\n")
	log.Println("GinLaravel is Working >>> ")
	log.Println("访问地址示例：http://" + serverAddr + "/gen1/app/api")

	// 启动服务
	err := httpServer.Run(serverAddr)

	if err != nil {
		panic("Run Error: " + err.Error())
	}
}
