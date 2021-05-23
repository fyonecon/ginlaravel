package Kit

import (
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
	"strings"
)

// Input 自动判断请求类型并自动输出参数值
func Input(ctx *gin.Context, key string) string {
	var _value string
	var value string
	var hasKey bool

	_method := ctx.Request.Method
	_contentType := ctx.Request.Header["Content-Type"]

	if _method == "GET" {
		value, hasKey = ctx.GetQuery(key)
	}else if _method == "POST" || _method == "OPTIONS" {
		value, hasKey = ctx.GetPostForm(key)

		if len(_contentType) >= 1 { // 判断是否含有请求头信息
			hasCt1 := strings.Contains(_contentType[0], "application/x-www-form-urlencoded")
			hasCt2 := strings.Contains(_contentType[0], "multipart/form-data")
			if !hasCt1 && !hasCt2 {
				_value = ""
				Common.Log("POST方式时建议：Content-Type=「application/x-www-form-urlencoded 」或 「 multipart/form-data 」")
			}else {
				_value = value
				Common.Log("当前请求头：" + _contentType[0])
			}
		}else {
			_value = ""
			Common.Log(_contentType)
		}

	}else {
		value, hasKey = "", false

	}

	if hasKey == false { // 参数不存在
		_value = ""
		// 当参数键不存在时，可能时是因为传来的参数的格式不正确。
		ctx.Request.ParseMultipartForm(128) //保存表单缓存的内存大小128M
		data := ctx.Request.Form
		//Common.Log("当参数键不存在时，可能时是因为传来的参数的格式不正确。请查看传来的GET+POST全部数据：")
		Common.Log(data)
		//Common.Log("axios请参考：项目资料/其他示例/axios-post.html")
	}else{
		_value = value
	}

	if key == "app_token" || key == "user_token" || key == "web_token" || key == "appToken" || key == "userToken" || key == "webToken" || key == "token" || key == "Token" { // 这些键不转义
		return _value
	}else {
		return Common.FilterInput(_value)
	}
}
