package Gen1App

import (
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

func Test1(ctx *gin.Context){

	method := ctx.Request.Method
	body := ctx.Request.Body
	header := ctx.Request.Header["Sec-Fetch-User"]

	id := Kit.Input(ctx, "id2") //

	// 接口返回
	ctx.JSON(200, gin.H{
		"method": method,
		"body": body,
		"header": header,
		"id": id,
	})
}