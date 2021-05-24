package Runtime
// http耗时服务参数

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
)

func StatLatency(ctx *gin.Context)  {
	StatLatency, _ := ctx.Get("stat_latency") // 获取ctx每次请求的全局值
	fmt.Println("\n 接口耗时=", StatLatency, "ms")

	_host, _ := ctx.Get("host")
	host := Common.ValueInterfaceToString(_host)
	Uri := host + ctx.Request.RequestURI
	fmt.Println(" 请求uri=", Uri)

	err := ctx.Request.ParseMultipartForm(128)
	if err != nil {
		fmt.Println("Request.Form报错=", err)
		ctx.Abort()
	} // 保存表单缓存的内存大小128M
	data := ctx.Request.Form
	fmt.Println(" 请求参数=", data)

	ctx.Next()
}

