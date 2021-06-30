package Gen3User

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context) {
	_timeMS := Common.GetTimeMS()
	timeMS := Common.IntToString(_timeMS)
	IP := ctx.ClientIP()
	where := "Gen3User"
	randString := Common.RandString(Common.RandRange(7, 9))

	_token := "" + IP + "#@" + timeMS + "#@" + where + "#@" + randString
	userToken := Kit.Encode(_token, "")

	back := map[string]interface{}{
		"user_id": "-1",
		"user_token": userToken,
		"time": Common.GetTimeDate("Y-m-d H:i:s"),
		"ip": IP,
	}

	// 接口返回
	ctx.JSONP(200, map[string]interface{}{
		"state": 1,
		"msg": "UserToken已生成",
		"content": back,
	})
}

