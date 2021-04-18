package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpCors(ctx *gin.Context) {
	method := ctx.Request.Method

	ctx.Header("Access-Control-Allow-Origin", "*")
	// ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}

	ctx.Next()
}

