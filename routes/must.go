package routes

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Http/Middleware"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Must 路由访问原则：宽进严出，所以都用Any，在拦截器里面拦截（VerifyXXX.go）具体请求事件。
// 路由周期：请求路由名——>header过滤——>拦截请求频率——>校验请求方法和Token参数——>运行目标函数——>程序达到终点，关闭此次请求。
// 路由写法 ：Any(路由名（必选）, header参数（可选）, 访问频率限制（可选）, 拦截器参数验证（可选）, 目标函数handler（必选）)
func Must(route *gin.Engine) { // ==系统必要路由==

	// 默认根路由
	route.Any("/", Middleware.HttpCorsApi, Middleware.HttpLimiter(2), func (ctx *gin.Context) {
		Kit.Log("进入了默认根路由", ctx.ClientIP())
		ctx.JSONP(http.StatusForbidden, gin.H{
			"state": 403,
			"msg": "指定路由名后才可访问",
			"content": map[string]interface{}{
				"gl_version": Common.ServerInfo["gl_version"],
				"go_version": Common.ServerInfo["go_version"],
				"timezone": Common.ServerInfo["timezone"],
			},
		})
	}, Middleware.HttpAbort)

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
	}, Middleware.HttpAbort)

	// ico图标
	route.StaticFile("/favicon.ico", Common.ServerInfo["go_path"] + "favicon.ico") // 单个静态文件

}
