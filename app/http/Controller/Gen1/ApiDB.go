package Gen1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiDB(ctx *gin.Context) {

	ctx.HTML(http.StatusOK,
		"index/index.html",
		gin.H{
			"title": "模版输出",
			"msg": "=index=",
		})

}

