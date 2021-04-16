package Gen1Controller

import (
	"ginlaravel/app/common"
	"ginlaravel/app/http/Model/Gen1Model"
	"github.com/gin-gonic/gin"
	"strconv"
)


func ListUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	_userClassId := ctx.Query("user_class_id")
	userClassId := common.StringToInt(_userClassId)
	_nickname := ctx.Query("nickname")
	// _page := ctx.PostForm("page")
	_page := ctx.Query("page")
	page := common.StringToInt(_page)
	page = page - 1

	listUserModel := Gen1Model.ListUserModel{}
	users, err := listUserModel.ListUser(page, userClassId, _nickname)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"
	}

	// 返回一些测试数据
	testData = map[string]string{
		"page": _page,
		"user_class_id": _userClassId,
		"nickname": _nickname,
	}

	ctx.JSON(200, gin.H{
		"state": state,
		"msg": msg,
		"test_data": testData,
		"content": users,
	})
}

// app/that_user?user_id=1
func ThatUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := ctx.Query("user_id") // GET方法获得参数
	userId, _ := strconv.Atoi(_userId)

	// 查询数据库
	userModel := Gen1Model.ThatUserModel{}
	data, err := userModel.ThatUser(userId)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"
	}

	// 返回一些测试数据
	testData = map[string]string{
		"user_id": _userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   data,
	})
}


func AddUser(ctx *gin.Context)  {

}


func EditUser(ctx *gin.Context)  {

}


func DelUser(ctx *gin.Context)  {

}
