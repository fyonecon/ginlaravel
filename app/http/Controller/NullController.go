package Controller

import (
	"github.com/gin-gonic/gin"
)

func Null(ctx *gin.Context) {
	ctx.String(403, "路由为空，请指名正确路由名")
}