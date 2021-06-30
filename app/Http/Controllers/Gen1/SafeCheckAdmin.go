package Gen1

import "github.com/gin-gonic/gin"

func VerifyAdmin(ctx *gin.Context) {
	ctx.Next()
}