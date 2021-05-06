package Middleware

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"ginlaravel/app/Common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HttpCorsApi 处理http-header信息
func HttpCorsApi(ctx *gin.Context) { // 面向Api
	method := ctx.Request.Method

	//
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//
	////
	ctx.Header("Content-type", "application/json; charset=utf-8")
	ctx.Header("Cache-Control", "max-age=0")
	ctx.Header("X-Powered-By", "GinLaravel; " + Common.ServerInfo["go_version"])
	ctx.Header("Author", "fyonecon")
	ctx.Header("Timezone", Common.ServerInfo["timezone"])
	ctx.Header("Date", Common.GetTimeDate("Y-m-d H:i:s"))
	//ctx.Header("Server", "Nginx")

	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	ctx.Next()
}

