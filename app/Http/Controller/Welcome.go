package Controller

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"github.com/gin-gonic/gin"
)

// Tpl
// @title 模版视图输出
// @description CorsWeb示例
// @Router / [GET]
func Tpl(ctx *gin.Context) {

	name := Kit.Input(ctx, "name")
	if len(name) == 0 { name = "name为空"}
	id := Kit.Input(ctx, "id")

	ctx.HTML(200, "pages/welcome/index.html", gin.H{
		"title": "Welcome GinLaravel !",
		"msg": "name=" + name + "；id=" + id,
	})
}

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

	ctx.JSON(200, back)
}

