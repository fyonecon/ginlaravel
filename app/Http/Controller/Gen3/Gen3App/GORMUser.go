package Gen3App

import (
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"ginlaravel/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var gDB *gorm.DB = driver.GDB

// ThatGUserKeys 某用户
type ThatGUserKeys struct { // 数据库键的结果集，需大写
	UserId int
	Nickname string
	CreatTime string
}
func ThatGUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")

	user := ThatUserKeys{} // 构建结果集

	// 多查询条件
	DBMap := map[string]interface{}{}
	DBMap["state"] = 1
	DBMap["user_id"] = userId

	// 操作数据库
	gDB.Table("gl_user").Where(DBMap).Find(&user)

	state = 1
	msg = "查询完成"

	// 访问结构体并改变成员变量的值
	createTime := user.CreatTime
	createTime = Common.DateToDate(createTime)
	user.CreatTime = createTime

	// 返回一些测试数据
	testData := gin.H{
		"user_id": userId,
		"DBMap": DBMap,
	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   user,
	})
}

