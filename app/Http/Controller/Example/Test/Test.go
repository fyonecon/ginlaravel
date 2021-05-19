package Test

import (
	"fmt"
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

func Test1(ctx *gin.Context) {

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

	//
	//Kit.MakeCaptcha(ctx, "123")
	//fmt.Println(imgCode)

	// 生成Excel
	rowData := [][]interface{}{
		{"name", "phone", "age"},
		{"张三1", "1231", 23},
		{"张三2", "1232", 24},
		{"张三3", "1233", 25},
		{"合计人数", 3},
	}
	_excelName := Common.MakeSMSCode(8)+".xlsx"
	excelName := Kit.MakeExcel(rowData, _excelName, "")
	excelState, _:= Common.HasFile(Common.ServerInfo["storage_path"] + "cache_file/"+excelName)

	// 读取Excel
	_rowData, _ := Kit.ReadExcel(excelName, "")
	fmt.Println(_rowData)

	//// 压缩图片示例
	//filepath := Common.ServerInfo["storage_path"] + "cache_file/" // 在服务器里面的绝对路径文件夹
	//dateFile := Common.GetTimeDate("Ymd") + "/"
	//saveFilepath := filepath + dateFile
	//// 创建日期文件夹
	//has, _ := Common.HasFile(saveFilepath)
	//if !has {
	//	err := os.Mkdir(saveFilepath, os.ModePerm)
	//	if err != nil {
	//		fmt.Printf("不能创建文件夹=[%v]\n", err)
	//	}
	//}
	//// 压缩图片
	//img := Kit.CompressImg("color.png", filepath, saveFilepath)

	// 接口返回
	ctx.JSON(200, gin.H{
		"method": method,
		"body": body,
		"header": header,
		"id": id,
		"nickname": nickname,
		//"img": img,
		"excel_name": excelName,
		"excel_state": excelState,
		////"timezone": Common.ServerInfo["timezone"],
		//"date": Common.GetTimeDate("Y-m-d H:i:s"),

	})
}
