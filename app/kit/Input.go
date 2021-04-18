package Kit

import "github.com/gin-gonic/gin"

// 自动判断请求类型并自动输出参数值
func Input(ctx *gin.Context, key string) string {
	var _value string
	var value string
	var hasKey bool

	_method := ctx.Request.Method

	if _method == "GET" {
		value, hasKey = ctx.GetQuery(key)
	}else if _method == "POST" {
		value, hasKey = ctx.GetPostForm(key)
	}else {
		value, hasKey = "(only GET/POST)", false
	}

	if !hasKey { // 参数不存在
		_value = ""
	}else{
		_value = value
	}

	return _value
}
