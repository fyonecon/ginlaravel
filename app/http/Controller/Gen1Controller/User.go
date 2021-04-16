package Gen1Controller

import (
	"ginlaravel/app/common"
	"ginlaravel/app/http/Model/Gen1Model"
	"github.com/gin-gonic/gin"
	"strconv"
)


func ListUser(ctx *gin.Context)  {
	state := "1"
	msg := "请求完成"

	_page := ctx.PostForm("page")
	page := common.StringToInt(_page)
	if page < 1 {page = 1}else if page > 200 {page = 200}

	_limit := ctx.PostForm("limit")
	limit := common.StringToInt(_limit)
	if limit < 1 {limit = 20}else if limit > 100 {limit = 100}

	// nickname := ctx.PostForm("nickname")

	res := ""

	ctx.JSON(200, gin.H{
		"state": state,
		"msg": msg,
		"content": res,
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
