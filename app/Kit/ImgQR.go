package Kit

import (
	"fmt"
	"ginvel.com/app/Common"
	qrCreate "github.com/skip2/go-qrcode"
	"image/color"
)

// 处理二维码

// MakeQR 生成二维码
// github.com/skip2/go-qrcode
func MakeQR(txt string) string {
	fmt.Println(txt)

	content := txt // 二维码内容
	width := 200 // 二维码宽=高，px
	filename := "color.png" // png，二维码绝对路径
	filepath := Common.ServerInfo["storage_path"] + "cache_file/" // 在服务器里面的绝对路径文件夹

	//带颜色的二维码
	if len(content) == 0 { content = "null" }
	if width < 80 { width = 80 }
	qr, _ := qrCreate.New(content, qrCreate.Medium)
	qr.DisableBorder = true
	qr.Level = 15 // 容错度，%
	qr.BackgroundColor = color.RGBA{R: 255, G: 255, B: 255, A: 255} // 背景色
	qr.ForegroundColor = color.RGBA{R:0, G:0, B:0, A:255} // 字色
	err := qr.WriteFile(width, filepath + filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return filename
}