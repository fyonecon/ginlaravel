package Gen1App

import (
	"ginlaravel/app/Common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TplIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/home/index.html", gin.H{
		"static_path": Common.ServerInfo["static_path"],
		"title": "模版输出-index",
		"msg": "=tpl=index=",
	})

}

func Tpl(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/test/test.html", gin.H{
		"static_path": Common.ServerInfo["static_path"],
		"title": "模版输出-test",
		"msg": "=tpl=test=",
	})

}