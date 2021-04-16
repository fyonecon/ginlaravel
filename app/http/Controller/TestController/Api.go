package TestController

import (
	"ginlaravel/app/common"
	"github.com/gin-gonic/gin"
)

// 接口输出简单数据
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

	// 引用公共函数和公共配置参数
	name3 := common.Test("common==func==")
	name4 := common.Config["test"]
	name5 := common.Config["api"]

	content := "name1==" + name1 + ", name2==" + name2 + ", name3=" + name3 + ", name4=" + name4+ ", name5=" + name5

	var back = map[string]string{
		"state": "1",
		"msg": "请求成功==",
		"content": content,
	}

	ctx.JSON(200, back)
}


// 接口输出复杂数据
type ArrayApi2 struct {
	Name   string
	Age    int64
	Info   map[string]string
}
func Api2(ctx *gin.Context)  {
	// 预定义接口解释类参数
	var state int
	var msg string

	var info = map[string]string{
		"job": "理发师",
		"avatar_url": "http://img",
	}

	// 构建多维数据
	back := ArrayApi2{
		Name:   "托尼老师",
		Age:    21,
		Info:   info,
	}

	state = 1
	msg = "请求成功"

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": state,
		"msg": msg,
		"content": back,
	})
}