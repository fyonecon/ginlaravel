package Kit

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// 自动判断请求类型并自动输出参数值
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
				log.Println("POST方式时建议：Content-Type=「application/x-www-form-urlencoded 」或 「 multipart/form-data 」")
			}else {
				_value = value
				log.Println("当前请求头：" + _contentType[0])
			}
		}else {
			_value = ""
			log.Println(_contentType)
		}

	}else {
		value, hasKey = "", false

	}

	if hasKey == false { // 参数不存在
		_value = ""
		// 当参数键不存在时，可能时是因为传来的参数的格式不正确。
		ctx.Request.ParseMultipartForm(128) //保存表单缓存的内存大小128M
		data := ctx.Request.Form
		log.Println("当参数键不存在时，可能时是因为传来的参数的格式不正确。请查看传来的GET+POST全部数据：")
		log.Println(data)
		log.Println("axios请参考：项目资料/其他示例/axios-post.html")
	}else{
		_value = value
	}

	return _value
}
