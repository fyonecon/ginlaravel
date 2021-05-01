package Gen3User

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

type UserLoginKeys struct {
	UserId string
	UserToken string
	Time string
	IP string
}
func UserLogin(ctx *gin.Context) {
	_timeMS := Common.GetTimeMS()
	timeMS := Common.IntToString(int(_timeMS))
	IP := ctx.ClientIP()
	where := "Gen3User"
	randString := Common.RandString(Common.RandRange(7, 9))

	_token := "" + IP + "#@" + timeMS + "#@" + where + "#@" + randString
	userToken := Kit.Encode(_token, "")

	back := UserLoginKeys{
		UserId: "-1",
		UserToken: userToken,
		Time: Common.GetTimeDate("Y-m-d H:i:s"),
		IP: IP,
	}

	// 接口返回
	ctx.JSONP(200, map[string]interface{}{
		"state": 1,
		"msg": "UserToken已生成",
		"content": back,
	})
}

