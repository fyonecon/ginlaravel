package Tpl

import (
	"ginlaravel/app/Common"
	"ginlaravel/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tpl1(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/test/test.html", gin.H{
		"title": "Welcome GinLaravel !",
		"msg": "=tpl=web=" + Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"],
	})

}

func Tpl2(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/test/test.html", gin.H{
		"title": "Welcome GinLaravel !",
		"msg": "=tpl=tpl=" + Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"],
	})

}