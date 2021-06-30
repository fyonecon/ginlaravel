package Gen1

import "github.com/gin-gonic/gin"

func VerifyApp(ctx *gin.Context) {
	ctx.Next()
}