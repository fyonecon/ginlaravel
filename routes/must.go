package routes

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Http/Middleware"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Must ==系统必要路由==
func Must(route *gin.Engine) {

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
