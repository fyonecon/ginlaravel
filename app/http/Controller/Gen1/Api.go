package Gen1

import (
	"github.com/gin-gonic/gin"
)

func Api(ctx *gin.Context) {

	// 请求GET参数
	name1 := ctx.Query("name1")
	name1 = string(name1)
	if len(name1) == 0 {
		name1 = "name1参数为空"
	}

	// 请求POST参数
	name2 := ctx.PostForm("name2")
	if len(name2) == 0 {
		name2 = "-1"
	}

	content := "name1=" + name1 + ", name2=" + name2

	var back = map[string]string{
		"state": "1",
		"msg": "请求成功",
		"content": content,
	}

	ctx.JSON(200, back)
}