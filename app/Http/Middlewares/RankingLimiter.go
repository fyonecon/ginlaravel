package Middlewares
// 路由优先级熔断式中间件

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
	"math"
)

// RankingLimiter 限制接口请求数据，做到主动降载，超阀值会自动熔断
// 优先级数字越大，优先级越低，规定：40代表cpu使用率超40%时拦截；95代表cpu使用率超过95%时拦截；250%代表cpu使用率超过250%时拦截。
// 数字最低是=40；最高=100*核心数（是否会因为多核超过100%没测过，暂时定为100%上限）；0表示不做任何熔断。
func RankingLimiter(ranking int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 请求链接信息
		// 读取每次请求的请求全局参数
		_host, _ := ctx.Get("host")
		host := Common.ValueInterfaceToString(_host)
		Uri := host + ctx.Request.RequestURI
		fmt.Println("请求uri=", Uri)

		// 读取超全局变量即可
		var _cpuNum interface{} = Common.GetGlobalData("cpu_num")
		var cpuNum int64 = _cpuNum.(int64)
		var _cpuPercent interface{} = Common.GetGlobalData("cpu_percent")
		var __cpuPercent float64 = _cpuPercent.(float64)
		var cpuPercent int64 = int64(math.Floor(__cpuPercent))

		var alertRanking string = "[80, 100]"
		//if ranking >= 80 && ranking <= 100*cpuNum { // 正确范围
		if ranking >= 80 && ranking <= 100 { // 正确范围
			ctx.Next()
		}else if ranking == 0 { // 0不做任何熔断操作
			ranking = 10000000
			ctx.Next()
		}else if ranking < 80 { // 使用默认值
			ranking = 80
			ctx.Next()
		} else {
			fmt.Println("优先级范围错误。" + alertRanking + " => ", cpuPercent)
			ctx.JSON(200, gin.H{
				"state": 0,
				"msg": "优先级范围错误",
				"content": gin.H{
					"Route-Ranking": ranking,
					"CPU-Num": cpuNum,
				},
			})

			ctx.Abort()
		}

		if cpuPercent > ranking { // 直接熔断，等待x秒定时周期后看CPU占用率是否恢复
			Kit.Log("达到熔断标准，uri=" + Uri, ctx.ClientIP())
			ctx.JSON(429, gin.H{
				"state": 429,
				"msg": "通道拥挤，请稍后再试",
				"content": gin.H{
					"Route-Ranking": ranking,
					"CPU-Percent": cpuPercent,
				},
			})
			ctx.Abort()
		}else {
			ctx.Next()
		}

	}
}
