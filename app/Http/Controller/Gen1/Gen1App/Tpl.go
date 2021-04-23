package Gen1App

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tpl(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "test/test.html", gin.H{
		"title": "模版输出",
		"msg": "=tpl=",
	})

}