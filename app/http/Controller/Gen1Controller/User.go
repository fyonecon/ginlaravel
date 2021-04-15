package Gen1Controller

import (
	"ginlaravel/app/common"
	"github.com/gin-gonic/gin"
)


func ListUser(ctx *gin.Context)  {
	state := "1"
	msg := "请求完成"

	_page := ctx.PostForm("page")
	page := common.StringToInt(_page)
	if page < 1 {page = 1}else if page > 200 {page = 200}

	_limit := ctx.PostForm("limit")
	limit := common.StringToInt(_limit)
	if limit < 1 {limit = 20}else if limit > 100 {limit = 100}

	// nickname := ctx.PostForm("nickname")

	res := ""

	ctx.JSON(200, gin.H{
		"state": state,
		"msg": msg,
		"content": res,
	})
}


func ThatUser(ctx *gin.Context)  {

}


func AddUser(ctx *gin.Context)  {

}


func EditUser(ctx *gin.Context)  {

}


func DelUser(ctx *gin.Context)  {

}
