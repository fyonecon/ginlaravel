package kit

import "github.com/gin-gonic/gin"

// 返回用户IP
func IP(ctx *gin.Context) string {
	return ctx.ClientIP()
}