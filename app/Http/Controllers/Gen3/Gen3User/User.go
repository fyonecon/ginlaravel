package Gen3User

import (
	"database/sql"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"ginvel.com/bootstrap/driver"
	"github.com/gin-gonic/gin"
	"log"
)

var DB *sql.DB = driver.MysqlDb // 连接gomysql扩展

// ListUserKeys 用户列表
type ListUserKeys struct { // 结果集，参数名需大写
	UserId int
	UserClassId int
	UserClassName string
	Nickname string
	CreatTime string
}
func ListUser(ctx *gin.Context)  {

	// 预定义参数
	var state int
	var msg string

	_page := Kit.Input(ctx, "page")
	_nickname := Kit.Input(ctx, "nickname")
	_userClassId := Kit.Input(ctx, "UserClassId")
	//_startTime := Kit.Input(ctx, "start_time")

	// 处理分页
	var limit int = Common.Page["limit"]
	var page int = int(Common.StringToInt(_page))
	var offset int = 0 // 本页从第几个开始
	if page <= 0 { page = 1 } else if page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 查询数据
	users := make([]ListUserKeys, 0) // 结果集

	// 多查询条件
	var DBString string // SQL语句
	var DBTotal string // 数据总数
	var DBMap string // where条件
	var DBOrder string // 排序
	var DBLimit string // 分页

	// where条件
	DBMap = " WHERE `state`=1 "
	if len(_userClassId) > 0 {
		DBMap =  DBMap + " AND `user_class_id`=" + _userClassId
	}
	if len(_nickname) > 0 {
		DBMap = DBMap + " AND `nickname` LIKE '%" + _nickname + "%'"
	}

	// 排序
	DBOrder = " ORDER BY `create_time` DESC, `nickname` ASC"
	// DBOrder = " ORDER BY `create_time` DESC"

	// 分页
	_offset := Common.IntToString(int64(offset))
	_limit := Common.IntToString(int64(limit))
	DBLimit = " LIMIT " + _offset + ", " + _limit

	// 拼装完整MySQL语句（注意查询语句顺序）
	DBString = "SELECT " +
		"`user_id`, `user_class_id`, `nickname`, `create_time` " +
		"FROM `gl_user` " +
		DBMap +
		DBOrder +
		DBLimit

	// 数据总数
	DBTotal =  "SELECT COUNT(`user_id`)" +
		"FROM `gl_user` " +
		DBMap

	// 查询
	rows, err := DB.Query(DBString + " ")
	totals, _ := DB.Query(DBTotal + " ")

	defer rows.Close()
	defer totals.Close()
	// 整理结果数组
	var user ListUserKeys
	for rows.Next() {
		rows.Scan(&user.UserId, &user.UserClassId, &user.Nickname, &user.CreatTime)
		users = append(users, user)
	}
	// 获取数据总数
	var total int
	for totals.Next() {
		totals.Scan(
			&total,
		)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"

		// 遍历切片中结构体，并改变结构体成员变量的值
		for i := 0; i < len(users); i++ {
			theCreateTime := users[i].CreatTime
			newCreateTime := Common.DateToDate(theCreateTime)
			users[i].CreatTime = newCreateTime
		}

	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"userClassId": _userClassId,
	}

	// 分页数据
	paging := map[string]interface{}{
		"total": total,
		"limit": limit,
		"page": page+1,
	}
	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"paging": paging,
		"content":   users,
	})
}

// ThatUserKeys 某用户信息
type ThatUserKeys struct { // 结果集，参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func ThatUser(ctx *gin.Context) {

	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")

	// 直接查询数据
	user := ThatUserKeys{} // 构建结果集
	err := DB.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`<>2 AND `user_id`=?", userId).Scan(&user.UserId, &user.Nickname, &user.CreatTime)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"

		// 访问结构体并改变成员变量的值
		createTime := user.CreatTime
		createTime = Common.DateToDate(createTime)
		user.CreatTime = createTime

	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"user_id": userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   user,
	})
}

// AddUser 新增用户信息
func AddUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userClassId := Kit.Input(ctx, "user_class_id")
	nickname := Kit.Input(ctx, "nickname")
	createTime := Common.GetTimeDate("YmhHis")

	data, err := DB.Exec("INSERT INTO `gl_user` (`user_class_id`, `nickname`, `create_time`) VALUES (?, ?, ?)", userClassId, nickname, createTime )

	_id, resErr := data.LastInsertId()

	if err != nil {
		state = 0
		msg = "新增失败"
	}else {
		state = 1
		msg = "新增成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"res_err": resErr,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   gin.H{
			"id": _id,
		},
	})
}

// UpdateUser 修改用户信息
func UpdateUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")

	userClassId := Kit.Input(ctx, "user_class_id")
	nickname := Kit.Input(ctx, "nickname")
	updateTime := Common.GetTimeDate("YmhHis")

	data, err := DB.Exec("UPDATE `gl_user` SET `user_class_id`=?, `nickname`=?, `update_time`=? WHERE `state`=1 AND `user_id`=? ", userClassId, nickname, updateTime, userId)

	res, resErr := data.RowsAffected()

	if err != nil {
		state = 0
		msg = "更新失败或无原数据"
	}else {
		state = 1
		msg = "更新成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"res_err": resErr,
		"user_id": userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   gin.H{
			"res": res,
		},
	})
}

// DelUser 删除用户
// 不是真正删除，只是不可见
func DelUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")
	updateTime := Common.GetTimeDate("YmhHis")

	data, err := DB.Exec("UPDATE `gl_user` SET `state`=2, `update_time`=? WHERE `state`=1 AND `user_id`=? ", updateTime, userId)

	res, resErr := data.RowsAffected()

	if err != nil {
		state = 0
		msg = "无原数据或已删除"
	}else {
		state = 1
		msg = "删除成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"res_err": resErr,
		"user_id": userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   gin.H{
			"res": res,
		},
	})
}

// ClearUser 清除用户
// 直接删除
func ClearUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string

	userId := Kit.Input(ctx, "user_id")

	data, err := DB.Exec("DELETE FROM `gl_user` WHERE `user_id` = ?", userId)

	res, resErr := data.RowsAffected()

	if err != nil {
		state = 0
		msg = "无原数据或已清除"
	}else {
		state = 1
		msg = "清除成功"
	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"res_err": resErr,
		"user_id": userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   gin.H{
			"res": res,
		},
	})
}
