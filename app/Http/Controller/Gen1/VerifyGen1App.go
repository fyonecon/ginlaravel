package Gen1

import "github.com/gin-gonic/gin"

func VerifyGen1App(ctx *gin.Context) {
	ctx.Next()
}