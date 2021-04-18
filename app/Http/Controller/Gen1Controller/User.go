package Gen1Controller

import (
	"ginlaravel/app/common"
	"ginlaravel/app/http/Model/Gen1Model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)


func ListUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	_userClassId := ctx.Query("user_class_id")
	userClassId := Common.StringToInt(_userClassId)
	_nickname := ctx.Query("nickname")
	// _page := ctx.PostForm("page")
	_page := ctx.Query("page")
	page := Common.StringToInt(_page)
	page = page - 1

	listUserModel := Gen1Model.ListUserKeys{}
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
	userModel := Gen1Model.ThatUserKeys{}
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

// 新增用户信息
func AddUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userClassId := ctx.Query("user_class_id") // GET方法获得参数
	userClassId := Common.StringToInt(_userClassId)
	_nickname := ctx.Query("nickname") // GET方法获得参数
	createTime := Common.GetTimeDate("YmdHis")

	if userClassId == 0 {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "user_class_id不正确",
		})
		return
	}

	if len(_nickname) == 0 {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "nickname不正确",
		})
		return
	}

	// 处理模型数据
	userModel := Gen1Model.AddUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		log.Println(err.Error())
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	userId, err := userModel.AddUser(int64(userClassId), _nickname, createTime)

	if err != nil {
		state = 0
		msg = "新增失败"
	}else {
		state = 1
		msg = "新增成功"
	}

	// 返回一些测试数据
	testData = map[string]string{

	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   userId,
	})
}

// 更新用户信息
func UpdateUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := ctx.Query("user_id") // GET方法获得参数
	userId := Common.StringToInt(_userId)

	_userClassId := ctx.Query("user_class_id") // GET方法获得参数
	userClassId := Common.StringToInt(_userClassId)
	_nickname := ctx.Query("nickname") // GET方法获得参数
	updateTime := Common.GetTimeDate("YmdHis")

	if userClassId == 0 {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "user_class_id不正确",
		})
		return
	}

	if len(_nickname) == 0 {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": "nickname不正确",
		})
		return
	}

	userModel := Gen1Model.UpdateUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	res, err := userModel.UpdateUser(int64(userId), int64(userClassId), _nickname, updateTime)

	if res == 0 || err != nil {
		// log.Println(err.Error())
		state = 0
		msg = "更新失败或用户原数据不存在"
	}else {
		state = 1
		msg = "更新成功"
	}

	// 返回一些测试数据
	testData = map[string]string{

	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   res,
	})

}

// 删除用户
// 不是真正删除
func DelUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := ctx.Query("user_id") // GET方法获得参数
	userId := Common.StringToInt(_userId)
	updateTime := Common.GetTimeDate("YmdHis")

	userModel := Gen1Model.DelUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	res, err := userModel.DelUser(int64(userId), updateTime)

	if res == 0 || err != nil {
		// log.Println(err.Error())
		state = 0
		msg = "用户数据不存在"
	}else {
		state = 1
		msg = "已删除"
	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"content":   res,
	})
}

// 彻底删除用户
func ClearUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := ctx.Query("user_id") // GET方法获得参数
	userId := Common.StringToInt(_userId)

	userModel := Gen1Model.ClearUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	res, err := userModel.ClearUser(int64(userId))

	if res == 0 || err != nil {
		// log.Println(err.Error())
		state = 0
		msg = "用户数据不存在"
	}else {
		state = 1
		msg = "已彻底删除"
	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"content":   res,
	})
}