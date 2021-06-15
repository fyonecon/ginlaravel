package Runtime
// http耗时服务参数

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
)

func StatLatency(ctx *gin.Context)  {

	_host, _ := ctx.Get("host")
	host := Common.ValueInterfaceToString(_host)
	Uri := host + ctx.Request.RequestURI
	fmt.Println("请求uri=", Uri)

	_statLatency, _ := ctx.Get("stat_latency") // 获取ctx每次请求的全局值
	fmt.Println("接口耗时=", _statLatency, "ms")

	//err := ctx.Request.ParseMultipartForm(128)
	//if err != nil {
	//	fmt.Println("Request.Form为空=", err)
	//	//ctx.Abort()
	//} // 保存表单缓存的内存大小128M
	//data := ctx.Request.Form
	//fmt.Println(" 请求参数=", data)

	statLatency := Common.StringToFloat(Common.ValueInterfaceToString(_statLatency)) // ms
	if statLatency > 3*1000 { // 超过3s都记录下来
		Kit.Log(Common.ValueInterfaceToString(_statLatency) + "ms；" + Uri, ctx.ClientIP())
	}

	ctx.Next()
}

