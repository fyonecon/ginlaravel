package routes
/*
路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

import (
	"fmt"
	"ginlaravel/app/Common"
	"ginlaravel/app/Http/Controller"
	"ginlaravel/app/Kit"
	"ginlaravel/app/Middleware"
	"ginlaravel/config"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//_ "ginlaravel/docs" // 跟目录执行「swag init」生成的docs文件夹路径，_引包表示只执行init函数。
	"github.com/gin-gonic/gin"
	"net/http"
)

// Must ==系统必要路由==
func Must(route *gin.Engine) {
	serverConfig := config.GetServerConfig()
	serverAddr := serverConfig["HOST"] + ":" + serverConfig["PORT"]
	fmt.Println(serverAddr)

	// 默认根路由
	route.Any("/", Middleware.HttpCorsWeb, Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		Kit.Log("进入了默认根路由", ctx.ClientIP())

		name := Kit.Input(ctx, "name")
		if len(name) == 0 { name = "name为空"}
		id := Kit.Input(ctx, "id")

		ctx.HTML(200, "pages/welcome/index.html", gin.H{
			"title": "Welcome GinLaravel !",
			"msg": "name=" + name + "；id=" + id,
		})

	})

	// 404路由
	route.NoRoute(Middleware.HttpCorsApi, Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var IP string = ctx.ClientIP()
		Kit.Error("404路由：" + url, IP)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 404,
			"msg": "未定义此名称的路由名",
			"content": map[string]interface{}{
				"url": url,
				"ip": IP,
			},
		})
	})

	// swagger接口文档，适配于GinLaravel
	//url := ginSwagger.URL("http://" + serverAddr + "/swagger/doc.json") // The url pointing to API definition
	//route.GET("/swagger/*any", Middleware.HttpCorsWeb, Middleware.HttpLimiter(2), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ico图标
	route.StaticFile("/favicon.ico", Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"] + "favicon.ico")

	// robots文件
	route.StaticFile("/robots.txt", Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"] + "robots.txt")

	// js、css、img等多个静态文件夹
	route.Static("/static/", Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"])

	// 示例-模版视图输出
	route.Any("tpl", Middleware.HttpCorsWeb, Middleware.HttpLimiter(2), Controller.Tpl)

	// 示例-api_json数据输出
	route.Any("api", Middleware.HttpCorsApi, Middleware.HttpLimiter(2), Controller.Api)

}
