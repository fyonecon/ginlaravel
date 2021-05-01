package Middleware
// 接口数据返回完成立即关闭下一步程序的等待，即接口运行达到终点，程序不必再等待。

import "github.com/gin-gonic/gin"

func HttpAbort(ctx *gin.Context) {
	ctx.Abort()
}
