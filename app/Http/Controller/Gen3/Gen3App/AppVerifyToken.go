package Gen3App
// 解析AppToken

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"strings"
)

func AppCheckToken(_appToken string) (int64, string, string) {
	var tokenTime int64
	var msg string

	_appToken, _ = Common.DecodeURL(_appToken) // 转义回来特殊字符
	_token := Kit.Decode(_appToken, "")
	if len(_token) == 0 { // token解析失败
		msg = "token解析失败"
		return 0, msg, _token
	}else { // token解析成功

		defer func() { // 跳过致命错误使程序继续运行
			if e := recover(); e != nil {
				tokenTime = 0 // 如果出现错误，直接让token过期即可
				msg = "token数据格式不正确"
			}
		}()

		_tokenArray := strings.Split(_token, "#@")
		// _ip := _tokenArray[0]
		_timeMS := _tokenArray[1]
		// _where := _tokenArray[2]

		tokenTime = int64(Common.StringToInt(_timeMS))
		nowTime := Common.GetTimeMS()

		theTime := nowTime - tokenTime
		if theTime < 1000*60*60*24*2 && theTime > 0 { // 有效
			msg = "token有效"
			return 1, msg, _token
		}else { // 过期
			msg = "token过期"
			return 0, msg, _token
		}
	}
}
