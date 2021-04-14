package Controller

import (
	"github.com/gin-gonic/gin"
)

func Null(ctx *gin.Context) {
	ctx.String(403, "路由为空")
}