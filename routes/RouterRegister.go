package routes

import (
	"ginvel.com/routes/RouterGroup"
	"github.com/gin-gonic/gin"
	"log"
)

// RouterRegister 灵活的注册路由文件
func RouterRegister(app *gin.Engine)  {
	log.Println("运行自定义注册路由文件 >>> ")
	// 示例
	RouterGroup.ApiExample(app) // 面向Api
	RouterGroup.WebExample(app) // 面向模版输出
	// 其他
	RouterGroup.ApiGen1(app)
	RouterGroup.ApiGen3(app)

	return
}