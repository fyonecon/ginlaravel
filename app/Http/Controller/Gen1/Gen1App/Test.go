package Gen1App

import (
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

func Test1(ctx *gin.Context){

	//fmt.Println("====1===")
	//
	//// x-www-form-urlencoded
	//ctx.Request.ParseForm()
	//for k, v := range ctx.Request.PostForm {
	//	fmt.Printf("k:%v\n", k)
	//	fmt.Printf("v:%v\n", v)
	//}
	//
	//fmt.Println("====2===")
	//
	//// form-data或x-www-form-urlencoded
	//ctx.Request.ParseMultipartForm(128) //保存表单缓存的内存大小128M
	//data := ctx.Request.Form
	//fmt.Println(data)
	//
	//fmt.Println("====2-1===")
	//
	method := ctx.Request.Method
	body := ctx.Request.Body
	header := ctx.Request.Header["Content-Type"]
	//
	//fmt.Println(ctx.Request.Form)
	//
	//fmt.Println("====3===")

	id := Kit.Input(ctx, "id") //
	nickname := Kit.Input(ctx, "nickname")

	// 接口返回
	ctx.JSON(200, gin.H{
		"method": method,
		"body": body,
		"header": header,
		"id": id,
		"nickname": nickname,
		//"timezone": Common.ServerInfo["timezone"],
		//"date": Common.GetTimeDate("Y-m-d H:i:s"),
	})
}