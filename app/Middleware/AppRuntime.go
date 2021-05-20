package Middleware
// 处理App运行时的一些必要事件

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Runtime"
	"ginlaravel/config"
	"github.com/gin-gonic/gin"
	"time"
)

// StatLatency 捕捉接口运行耗时
// 此处使用的Next()请参考文档：https://blog.csdn.net/qq_37767455/article/details/104712028
func StatLatency(ctx *gin.Context) {
	start := float64(time.Now().UnixNano()) / 1000000 // ms
	ctx.Set("stat_start", Common.GetTimeDate("i.s.ms.ns")) // 设置公共参数

	// 等其他中间件先执行
	ctx.Next()

	// 获取运行耗时，ms
	end := float64(time.Now().UnixNano()) / 1000000
	latency := end - start

	// 设置公共参数
	ctx.Set("stat_latency", latency)
	//fmt.Println("本次运行耗时=", latency, "ms")

	// 进入Http服务自我治理
	Runtime.HttpServer(ctx)

	// 进入硬件自我治理服务
	Runtime.Hardware(ctx)

	// 计时完成，中断所有后续函数调用
	ctx.Abort()
}

// AppData 设置全局参数
func AppData(ctx *gin.Context) {
	serverConfig := config.GetServerConfig()
	host := serverConfig["HOST"] + ":" + serverConfig["PORT"]

	ctx.Set("host", host)

	ctx.Next()
}