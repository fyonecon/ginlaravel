package Controller
// 接口访问拦截器

import "github.com/gin-gonic/gin"

func VerifyGen1(ctx *gin.Context) {
	ctx.Next() // 检测通过，继续下一步操作
}
