package Captcha
// 返回图形验证码

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)


func Captcha(ctx *gin.Context)  {
	code := Common.MakeSMSCode(6) // 随机数字
	Kit.MakeCaptcha(ctx, code)
}
