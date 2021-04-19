package Controller
// 接口访问拦截器

import (
	"github.com/gin-gonic/gin"
)

// 接口访问校验，校验不过则立即中断访问
func VerifyTest(ctx *gin.Context) {
	// 校验接口访问方式
	method := ctx.Request.Method
	if method == "GET" {
		ctx.Next() // 检测通过，继续下一步操作
	}else if method == "POST" || method == "OPTION" {
		ctx.Next() // 检测通过，继续下一步操作
	}else {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "未知的访问方法",
			"content": map[string]string{
				"method": method,
				"IP": ctx.ClientIP(),
			},
		})
		ctx.Abort() // 中断下一步函数运用
	}
}
