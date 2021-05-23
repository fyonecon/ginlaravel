package Runtime
// http服务参数

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
)

func HttpServer(ctx *gin.Context)  {
	StatLatency, _ := ctx.Get("stat_latency")
	fmt.Println("\n 接口耗时=", StatLatency, "ms")

	_host, _ := ctx.Get("host")
	host := Common.ValueInterfaceToString(_host)
	Uri := host + ctx.Request.RequestURI
	fmt.Println(" 请求uri=", Uri)

	ctx.Request.ParseMultipartForm(128) // 保存表单缓存的内存大小128M
	data := ctx.Request.Form
	fmt.Println(" 请求参数=", data)

	ctx.Next()
}

