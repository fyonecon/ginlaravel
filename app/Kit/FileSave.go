package Kit

import (
	"bufio"
	"fmt"
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

// DownloadImg 下载url图片，不压缩
func DownloadImg(imgUrl string) (imgName string, saveFilepath string) {
	filepath := Common.ServerInfo["storage_path"]+"upload/"

	// 创建文件夹
	dateFile := Common.GetTimeDate("dl-Ymd") + "/"
	saveFilepath = filepath + dateFile
	// 创建日期文件夹
	has, _ := Common.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹2=[%v]\n", err)
		}
	}

	fileName := path.Base(imgUrl)
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("不能下载图片")
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024)

	file, err := os.Create(saveFilepath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)

	fmt.Printf("Total length: %d", written)

	imgName = fileName
	return
}

// SaveFile 保存文件（图片、pdf、zip等）
// file, _ := ctx.FormFile("file") // file是表单字段名字，如<input type="file" name="file">
// 调用：filename := Kit.SaveFile(ctx, file)
func SaveFile(ctx *gin.Context, file *multipart.FileHeader) (newFilename string) {
	filepath := Common.ServerInfo["storage_path"]+"upload/"

	filename := file.Filename // 文件名
	fileSize := file.Size // 文件大小

	// 打印上传的文件名
	fmt.Println(filename, fileSize)

	// 创建文件夹
	dateFile := Common.GetTimeDate("Save-Ymd") + "/"
	saveFilepath := filepath + dateFile
	// 创建日期文件夹
	has, _ := Common.HasFile(saveFilepath)
	if !has {
		err := os.Mkdir(saveFilepath, os.ModePerm)
		if err != nil {
			fmt.Printf("不能创建文件夹1=[%v]\n", err)
		}
	}

	_filename := filename + "_" + newFileName()

	err := ctx.SaveUploadedFile(file, filepath +_filename)
	if err != nil {
		return ""
	}

	newFilename = _filename
	return
}


// 重新起名
func newFileName() string {
	return fmt.Sprintf(Common.GetTimeDate("YmdHis")+"_"+Common.RandString(Common.RandRange(6, 8))+"")
}