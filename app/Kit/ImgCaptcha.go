package Kit

import (
	"ginvel.com/app/Common"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
)

// MakeCaptcha 生成图形验证码
// _code自定义的图片显示的文字
// 调用：Kit.MakeCaptcha(ctx, "123456") ，并且不需要任何return或ctx.JSON()
// 返回image/png
func MakeCaptcha(ctx *gin.Context, _code string) *captcha.Image {

	if len(_code) < 4 || len(_code) > 8 {
		_code = "null"
	}

	// 图片宽高
	height := len(_code)*10
	width := height*2-height/5

	w := ctx.Writer // 等于 w http.ResponseWriter
	// r := ctx.Request // 等于 r *http.Request

	cap := captcha.New()
	// 可以设置多个扰乱字体 或使用cap.AddFont("xx.ttf")追加
	cap.SetFont(Common.ServerInfo["storage_path"] + "sys_file/afocus/comic.ttf")
	// 设置验证码大小，px
	cap.SetSize(width, height) // 128,64
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{R: 255, G: 255, B: 255, A: 200})
	// 设置背景色 可以多个 随机替换背景色 默认白色
	cap.SetBkgColor(
		color.RGBA{R:255, G:0, B:0, A:50},
		color.RGBA{R:0, G:0, B:255, A:80},
		color.RGBA{R:0, G:153, B:0, A:110},
	)

	w.Header().Set("Content-Type", "image/png; charset=utf-8") // 强制开启返回给浏览器的格式

	pixImg := cap.CreateCustom(_code) // 自定义的验证码
	// pixImg, code := cap.Create(6, captcha.ALL) // 使用系统生成的验证码
	// return pixImg, code

	err := png.Encode(w, pixImg)
	if err != nil {
		return nil
	} // 将pix格式转换成png

	return pixImg
}