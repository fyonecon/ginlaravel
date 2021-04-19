package Gen3

import "github.com/gin-gonic/gin"

func VerifyOpen(ctx *gin.Context) {
	ctx.Next()
}