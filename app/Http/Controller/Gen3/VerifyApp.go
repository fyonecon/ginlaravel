package Gen3

import (
	"ginlaravel/app/Http/Controller/Gen3/Gen3App"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

func VerifyApp(ctx *gin.Context) {
	IP := ctx.ClientIP()

	// 校验接口访问方式
	method := ctx.Request.Method
	if method == "GET" {
		//ctx.JSON(200, gin.H{
		//	"state": 0,
		//	"msg": "不允许使用GET方法请求数据",
		//	"content": map[string]string{
		//		"method": method,
		//		"IP": ctx.ClientIP(),
		//	},
		//})
		//ctx.Abort() // 中断下一步函数运用

		// 校验token
		appToken := Kit.Input(ctx, "AppToken")
		_state, _msg, _ := Gen3App.AppCheckToken(appToken)
		if _state == 1 {
			ctx.Next() // 检测通过，继续下一步操作

		}else {
			if len(_msg) == 0 {
				_msg = "token数据格式不正确"
				Kit.Log(_msg, IP)
			}
			ctx.JSON(200, gin.H{
				"state": _state,
				"msg": "" + _msg,
				"content": appToken,
			})
			ctx.Abort()
		}

	}else if method == "POST" || method == "OPTION" {
		// 校验token
		appToken := Kit.Input(ctx, "AppToken")
		_state, _msg, _ := Gen3App.AppCheckToken(appToken)
		if _state == 1 {
			ctx.Next() // 检测通过，继续下一步操作

		}else {
			if len(_msg) == 0 {
				_msg = "token数据格式不正确"
				Kit.Log(_msg, IP)
			}
			ctx.JSON(200, gin.H{
				"state": _state,
				"msg": "" + _msg,
				"content": appToken,
			})
			ctx.Abort()
		}
	}else {
		Kit.Log(method, IP)
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "未知的访问方法",
			"content": map[string]string{
				"method": method,
				"IP": IP,
			},
		})
		ctx.Abort() // 中断下一步函数运用
	}

}