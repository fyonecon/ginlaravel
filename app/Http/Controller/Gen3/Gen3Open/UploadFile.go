package Gen3Open

import (
	"ginlaravel/app/Common"
	"github.com/gin-gonic/gin"
	"log"
)

// Form法上传文件
// <form action="http://127.0.0.1:8090/test/upload_form_file" method="post" enctype="multipart/form-data">
//    文件: <input type="file" name="file"><br><br>
//    <input type="submit" value="上传文件">
//</form>
func UploadFormFile(ctx *gin.Context)  {
	defer func() { // 跳过致命错误使程序继续运行
		if e := recover(); e != nil {
			log.Println("文件上传接口报错。。")
			// 返回特殊格式意义的数据
			ctx.JSON(200, gin.H{
				"state": 0,
				"msg": "文件上传接口报错。。",
				"content": e,
			})
		}
	}()
	// 获取上传文件，返回的是multipart.FileHeader对象，代表一个文件，里面包含了文件名之类的详细信息
	// file是表单字段名字
	file, _ := ctx.FormFile("file")
	// 文件名
	filename := file.Filename
	// 文件格式
	fileSize := file.Size

	// 打印上传的文件名
	log.Println(filename)
	log.Println(fileSize)

	_filename := "2021-04_" + Common.RandString(Common.RandRange(5, 7)) + "_" + filename
	ctx.SaveUploadedFile(file, Common.ServerInfo["storage_path"] + "upload/" + _filename)

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "文件上传完成",
		"content": _filename,
	})

	// ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! Size：%d", filename, fileSize))
	// '截屏2021-04-20 下午3.34.33.png' uploaded! Size：459743

}
