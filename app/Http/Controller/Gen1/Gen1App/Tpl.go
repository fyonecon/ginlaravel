package Gen1App

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TplIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "模版输出-index",
		"msg": "=tpl=index=",
	})

}

func Tpl(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "test/test.html", gin.H{
		"title": "模版输出-test",
		"msg": "=tpl=test=",
	})

}