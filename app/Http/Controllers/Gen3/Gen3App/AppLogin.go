package Gen3App

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
)

func GetAppToken(ctx *gin.Context) {
	_timeMS := Common.GetTimeMS()
	timeMS := Common.IntToString(_timeMS)
	IP := ctx.ClientIP()
	where := "Gen3App"
	randString := Common.RandString(Common.RandRange(4, 6))

	_token := "" + IP + "#@" + timeMS + "#@" + where + "#@" + randString
	appToken := Kit.Encode(_token, "")
	appToken = Common.EncodeURL(appToken) // 顺便转义一下特殊字符，让appToken也可在GET方法中使用

	back := map[string]interface{}{
		"app_token": appToken,
		"time": Common.GetTimeDate("Y-m-d H:i:s"),
		"ip": IP,
	}

	// 接口返回
	ctx.JSONP(200, map[string]interface{}{
		"state": 1,
		"msg": "AppToken已生成",
		"content": back,
	})
}
