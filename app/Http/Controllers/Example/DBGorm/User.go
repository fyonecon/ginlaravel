package DBGorm

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
)

// ThatGUser 某用户
func ThatGUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")

	// 数据库表的字段（输出结果）
	type ResKeys struct {
		UserId int `json:"user_id"`
		Nickname string `json:"nickname"`
		UserClassId string `json:"user_class_id"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
	}

	// 多查询条件
	WhereMap := map[string]interface{}{}
	WhereMap["state"] = 1
	WhereMap["user_id"] = userId

	// 操作数据库
	res := ResKeys{}
	Kit.DB.Table("gl_user").Where(WhereMap).Find(&res)

	if res.UserId == 0 {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"
	}

	// 访问结构体并改变成员变量的值
	createTime := res.CreateTime
	createTime = Common.DateToDate(createTime)
	res.CreateTime = createTime

	// 返回一些测试数据
	testData := map[string]interface{}{
		"user_id": userId,
		"WhereMap": WhereMap,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   res,
	})

}



