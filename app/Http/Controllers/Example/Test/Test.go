package Test

import (
	"bufio"
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

func Test2(ctx *gin.Context)  {

	oldFilepath := Common.ServerInfo["storage_path"] + "cache_file/"
	newFilepath := Common.ServerInfo["storage_path"] + "cache_file/"
	filename := "phone.txt"

	// 计数
	lineNumHave := 0
	lineNumNull := 0

	// 打开txt文件
	file, err := os.Open(oldFilepath + filename)
	if err != nil{
		fmt.Println("无效的txt文件", err)
		return
		//os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// 创建新txt文件
	newName := "new_" + Common.GetTimeDate("YmdHis") + "_" + filename
	newFile, err := os.OpenFile(newFilepath + newName, os.O_CREATE | os.O_APPEND |os.O_WRONLY, 0666)
	if err != nil {
		log.Println("文件创建失败", newFilepath, err)
		//panic(err)
		return
	}
	defer newFile.Close()

	// 按行处理txt
	for scanner.Scan(){
		//fmt.Println(strings.TrimSpace(scanner.Text()))
		lineTxt := strings.TrimSpace(scanner.Text())

		// 写入文件内容
		if len(lineTxt) == 0 {
			_, _ = newFile.WriteString(lineTxt + "\n")
			lineNumNull ++
		}else {
			lineTxt = Common.MD5("86" + lineTxt)
			_, _ = newFile.WriteString(lineTxt + "\n")
			lineNumHave ++
		}

	}

	msgNewFile := "\n\n 有效值=" + Common.IntToString(int64(lineNumHave)) + "；空值=" + Common.IntToString(int64(lineNumNull)) + "\n\n"
	_, _ = newFile.WriteString(msgNewFile)

	// 接口返回
	back := gin.H{
		"newName": newName,
		"lineNumHave": lineNumHave,
		"lineNumNull": lineNumNull,
	}
	ctx.JSON(200, back)
}

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

	//// 生成Excel
	//rowData := [][]interface{}{
	//	{"name", "phone", "age"},
	//	{"张三1", "1231", 23},
	//	{"张三2", "1232", 24},
	//	{"张三3", "1233", 25},
	//	{"合计人数", 3},
	//}
	//_excelName := Common.MakeSMSCode(8)+".xlsx"
	//excelName := Kit.MakeExcel(rowData, _excelName, "")
	//excelState, _:= Common.HasFile(Common.ServerInfo["storage_path"] + "cache_file/"+excelName)
	//
	//// 读取Excel
	//_rowData, _ := Kit.ReadExcel(excelName, "")
	//fmt.Println(_rowData)

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

	// 获取进程启动时间
	//var stat os.FileInfo
	//if stat, _ = os.Lstat(fmt.Sprintf("/proc/%v", pid)); err != nil {
	//	fmt.Println(err)
	//}
	//proc.mtime = stat.ModTime().Unix()

	// 接口返回
	back := gin.H{
		"method": method,
		"body": body,
		"header": header,
		"id": id,
		"nickname": nickname,
		//"latency": ctx.Get("state_latency"),
		//"img": img,
		//"excel_name": excelName,
		//"excel_state": excelState,
		////"timezone": Common.ServerInfo["timezone"],
		//"date": Common.GetTimeDate("Y-m-d H:i:s"),

	}
	ctx.JSON(200, back)
}
