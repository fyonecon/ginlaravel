package Middleware

/*
第二作者Author：fyonecon
博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
Github：https://github.com/fyonecon/ginlaravel
邮箱Email：2652335796@qq.com，ikydee@yahoo.com
微信WeChat：fy66881159
所在城市City：长沙ChangSha
*/

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"time"
)

// 拦截http请求频率
func HttpLimiter(_max float64) gin.HandlerFunc {
	var max float64
	var tbOptions limiter.ExpirableOptions
	tbOptions.DefaultExpirationTTL = time.Second // 默认按每秒

	if _max > 0 || _max <= 200 {
		max = _max
	}else {
		max = 4
	}
	lmt := tollbooth.NewLimiter(max, &tbOptions) // 默认4次/秒，建议范围[1，20]

	return func(ctx *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if httpError != nil {
			//ctx.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			ctx.JSON(429, gin.H{
				"state": 429, "msg": "访问频率限制", "content": "",
			})

			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
