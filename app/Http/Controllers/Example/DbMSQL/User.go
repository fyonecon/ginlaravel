package DbMSQL

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"github.com/gin-gonic/gin"
	"log"
)

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
	rows, err := Kit.Db.Query(DBString + " ")
	totals, _ := Kit.Db.Query(DBTotal + " ")

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

// ThatUserKeys 某用户
type ThatUserKeys struct { // 结果集，参数名需大写
	UserId int
	Nickname string
	CreateTime string
}
func ThatUser(ctx *gin.Context) {

	// 预定义参数
	var state int
	var msg string

	_userId := Kit.Input(ctx, "user_id")
	userId := Common.StringToInt(_userId)

	// 直接查询数据
	user := ThatUserKeys{} // 构建结果集
	err := Kit.Db.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 AND `user_id`=?", userId).Scan(&user.UserId, &user.Nickname, &user.CreateTime)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"

		// 访问结构体并改变成员变量的值
		createTime := user.CreateTime
		createTime = Common.DateToDate(createTime)
		user.CreateTime = createTime

	}

	// 返回一些测试数据
	testData := map[string]interface{}{
		"user_id": _userId,
	}

	// 返回特殊格式意义的数据
	ctx.JSONP(200, map[string]interface{}{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"content":   user,
	})
}
