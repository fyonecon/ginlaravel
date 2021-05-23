package Kit
// 对称加密算法
// 调用：Encode()、Decode()
// 加密或解密失败，会返回空字符串

import (
	"ginvel.com/app/Kit/KitLib"
)

// Encode 加密
func Encode(text string, key string) string {
	sec := KitLib.Secret{}
	return sec.Encode(text, key)
}

// Decode 解密
func Decode(text string, key string) string {
	sec := KitLib.Secret{}
	return sec.Decode(text, key)
}
