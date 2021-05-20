package Controller

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

// Api
// @title 接口输出简单数据
// @description CorsApi示例
// @Router / [GET]
func Api(ctx *gin.Context) {

	name := Kit.Input(ctx, "name")
	if len(name) == 0 { name = "name为空"}
	_id := Kit.Input(ctx, "id")
	id := Common.StringToInt(_id)

	content := map[string]interface{}{
		"name": name,
		"id": id,
	}

	var back = map[string]interface{}{
		"state": 1,
		"msg": "接口请求成功",
		"content": content,
	}

	ctx.JSONP(200, back)
}

