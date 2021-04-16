package door

import (
	"fmt"
	"ginlaravel/config"
	"ginlaravel/routes"
	"github.com/gin-gonic/gin"
)

// 配置并启动http服务
// 项目访问入口
func Run(httpServer *gin.Engine) {
	fmt.Println("检查中间件...")

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

	// 启动服务
	err := httpServer.Run(serverAddr)

	if nil != err {
		panic("server run error: " + err.Error())
	}

}
