package Tpl

import (
	"ginvel.com/app/Common"
	"ginvel.com/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Tpl1(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/test/test.html", gin.H{
		"title": "Welcome ginvel.com !",
		"msg": "=tpl=web=" + Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"],
	})

}

func Tpl2(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "pages/test/test.html", gin.H{
		"title": "Welcome ginvel.com !",
		"msg": "=tpl=tpl=" + Common.ServerInfo["go_path"] + config.GetViewConfig()["View_Static"],
	})

}