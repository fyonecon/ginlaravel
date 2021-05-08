package Example
// 拦截器

import "github.com/gin-gonic/gin"

func VerifyExample(ctx *gin.Context) {
	ctx.Next()
}