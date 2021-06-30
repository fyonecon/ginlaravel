package routes

/*
路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
路由命名原则：推荐使用4层路由。第1层：api类还是web类；第2层：接口版本名；第3层：不同拦截器下的不同空间命名；第4层：目标函数handler。
*/

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Http/Controllers"
	"ginvel.com/app/Http/Middlewares"
	"ginvel.com/app/Kit"
	"ginvel.com/config"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//_ "GinLaravel/docs" // 跟目录执行「swag init」生成的docs文件夹路径，_引包表示只执行init函数。
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouteMust ==系统必要路由==
// handlerChain含义：
// 必选：Middlewares.HttpCorsApi 输出json-header
// 可选：Middlewares.HttpLimiter(10) 限流，数字代表每秒最多访问多少次就会禁止访问
// 可选：Middlewares.RankingLimiter(100) 熔断，数字代表CPU占用率到到xxx%就会熔断此接口
// 必选：Controllers.Welcome 目标函数，目标控制器到函数
func RouteMust(route *gin.Engine) {
	var staticPath string = Common.ServerInfo["framework_path"] + config.GetViewConfig()["View_Static"] // 静态文件目录

	// 默认根路由
	route.Any("/", Middlewares.HttpCorsApi, Middlewares.HttpLimiter(2), Middlewares.RankingLimiter(100), Controllers.Welcome)

	// 404路由
	route.NoRoute(Middlewares.HttpCorsApi, Middlewares.HttpLimiter(2), func (ctx *gin.Context) {
		var url string = ctx.Request.Host + ctx.Request.URL.Path
		var IP string = ctx.ClientIP()
		Kit.Error("404路由 >>> " + url, IP)
		ctx.JSONP(http.StatusNotFound, gin.H{
			"state": 404,
			"msg": "GinLaravel：未定义此名称的路由名",
			"content": map[string]interface{}{
				"url": url,
				"time": Common.GetTimeDate("Ymd.His.ms.ns"),
			},
		})
	})

	// swagger接口文档，适配于GinLaravel
	//url := ginSwagger.URL("http://" + serverAddr + "/swagger/doc.json") // The url pointing to API definition
	//route.GET("/swagger/*any", HttpCorsWeb, HttpLimiter(2), ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ico图标
	route.StaticFile("/favicon.ico", staticPath + "favicon.ico")
	// robots文件
	route.StaticFile("/robots.txt", staticPath + "robots.txt")

	// 示例-api_json数据输出
	route.Any("api", Middlewares.HttpCorsApi, Middlewares.HttpLimiter(200), Middlewares.RankingLimiter(100), Controllers.Api)

}
