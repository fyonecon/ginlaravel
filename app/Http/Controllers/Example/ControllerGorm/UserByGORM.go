package ControllerGorm

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Http/Models/Example/ModelGorm"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
	"log"
)

// ListUser 用户列表
func ListUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	type ListUserKeys struct { // 结果集，参数名需大写
		UserId int
		UserClassId int
		UserClassName string
		Nickname string
		CreatTime string
	}

	_userClassId := Kit.Input(ctx, "user_class_id")
	userClassId := Common.StringToInt(_userClassId)
	nickname := Kit.Input(ctx, "nickname")

	_page := Kit.Input(ctx, "page")
	page := int(Common.StringToInt(_page))

	// 处理分页
	var limit int = Common.Page["limit"]
	var offset int = 0 // 本页从第几个开始
	if page <= 0 { page = 1 } else if page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 将多查询条件转换成string
	keyArray := [...]interface{}{userClassId, nickname, page}
	keyString := Common.MakeRedisKey(keyArray)

	// 查询redis中是否存在该键
	hasValue := Kit.GetCacheInput(ctx, keyString)
	if len(hasValue) != 0 {
		fmt.Println(hasValue)
		hasValue["msg"] = "获取缓存数据成功"
		ctx.JSONP(200, hasValue)
	}else {

		listUserModel := ModelGorm.ListUserKeys{}
		res, total, err := listUserModel.ListUser(limit, offset, userClassId, nickname)

		if err != nil {
			state = 0
			msg = "查询无数据"
		}else {
			state = 1
			msg = "查询完成"
		}

		// 返回一些测试数据
		testData := map[string]interface{}{
			"page": _page,
			"user_class_id": userClassId,
			"nickname": nickname,
		}

		// 分页数据
		paging := Common.MakePaging(int(total), limit, page)
		// 返回数据
		back := map[string]interface{}{
			"state": state,
			"msg": msg,
			"paging": paging,
			"test_data": testData,
			"content": res,
		}
		cc := Kit.CreateCacheInput(ctx, keyString, back)
		fmt.Println(cc)

		ctx.JSONP(200, back)
	}

}

// ThatUser 某用户
func ThatUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	_userId := Kit.Input(ctx, "user_id")
	userId := Common.StringToInt(_userId)

	// 查询数据库
	userModel := ModelGorm.ThatUserKeys{}
	data, err := userModel.ThatUser(userId)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
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

// AddUser 新增用户信息
func AddUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	_userClassId := Kit.Input(ctx, "user_class_id")
	userClassId := Common.StringToInt(_userClassId)
	_nickname := Kit.Input(ctx, "nickname")
	createTime := Common.GetTimeDate("YmdHis")

	if userClassId == 0 {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "user_class_id不正确",
		})
		return
	}

	if len(_nickname) == 0 {
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": "nickname不正确",
		})
		return
	}

	// 处理模型数据
	userModel := ModelGorm.AddUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		log.Println(err.Error())
		ctx.JSONP(200, map[string]interface{}{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	userId, err := userModel.AddUser(userClassId, _nickname, createTime)

	if err != nil {
		state = 0
		msg = "新增失败"
	}else {
		state = 1
		msg = "新增成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{

	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   userId,
	})
}

// UpdateUser 更新用户信息
func UpdateUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := Kit.Input(ctx, "user_id")
	userId := Common.StringToInt(_userId)
	_userClassId := Kit.Input(ctx, "user_class_id")
	userClassId := Common.StringToInt(_userClassId)
	_nickname := Kit.Input(ctx, "nickname") // GET方法获得参数
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

	userModel := ModelGorm.UpdateUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	res, err := userModel.UpdateUser(userId, userClassId, _nickname, updateTime)

	if res == 0 || err != nil {
		// log.Println(err.Error())
		state = 0
		msg = "更新失败或用户原数据不存在"
	}else {
		state = 1
		msg = "更新成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{

	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   res,
	})

}

// DelUser 删除用户
// 不是真正删除
func DelUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	// _userId := ctx.PostForm("user_id") // Post方法获得参数
	_userId := Kit.Input(ctx, "user_id") // GET方法获得参数
	userId := Common.StringToInt(_userId)
	updateTime := Common.GetTimeDate("YmdHis")

	userModel := ModelGorm.DelUserKeys{}
	if err := ctx.ShouldBind(&userModel); err != nil {
		ctx.JSON(200, gin.H{
			"state": 0,
			"msg": err.Error(),
		})
		return
	}

	// 操作数据
	res, err := userModel.DelUser(userId, updateTime)

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

// ClearUser 彻底删除用户
func ClearUser(ctx *gin.Context)  {
	// 预定义参数
	var state int
	var msg string

	// 处理请求参数
	_userId := Kit.Input(ctx, "user_id") // GET方法获得参数
	userId := Common.StringToInt(_userId)

	userModel := ModelGorm.ClearUserKeys{}
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