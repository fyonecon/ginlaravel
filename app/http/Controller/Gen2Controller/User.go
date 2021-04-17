package Gen2Controller

import (
	"database/sql"
	"ginlaravel/app/common"
	"ginlaravel/app/http/Controller"
	"ginlaravel/app/kit"
	"ginlaravel/app/provider/driver"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

var DB *sql.DB = driver.MysqlDb

// 用户列表
type ListUserKeys struct { // 结果集，参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func ListUser(ctx *gin.Context)  {
	Controller.Gen2SafeCheck(ctx) // 检测请求是否安全

	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	_page := kit.Input(ctx, "page")
	_nickname := kit.Input(ctx, "nickname")

	// 处理分页
	var limit int = common.Page["limit"]
	var page int = common.StringToInt(_page)
	var offset int = 0 // 本页从第几个开始
	if page <= 0 { page = 1 } else if page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 构建查询
	nickname := "%" + _nickname + "%" // 模糊查询

	// 查询数据
	users := make([]ListUserKeys, 0) // 结果集
	rows, err := DB.Query("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 AND `nickname` LIKE ? LIMIT ?, ?", nickname, offset, limit)
	defer rows.Close()
	// 整理结果集
	var user ListUserKeys
	for rows.Next() {
		rows.Scan(&user.UserId, &user.Nickname, &user.CreatTime)
		users = append(users, user)
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
			newCreateTime := common.DateToDate(theCreateTime)
			users[i].CreatTime = newCreateTime
		}

	}

	// 返回一些测试数据
	testData = map[string]string{

	}

	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"state":     state,
		"msg":       msg,
		"test_data": testData,
		"users_type": reflect.TypeOf(users),
		"content":   users,
	})
}

// 某用户
type ThatUserKeys struct { // 结果集，参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func ThatUser(ctx *gin.Context) {
	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	_userId := kit.Input(ctx, "user_id")
	userId := common.StringToInt(_userId)

	// 直接查询数据
	user := ThatUserKeys{} // 构建结果集
	err := DB.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 AND `user_id`=?", userId).Scan(&user.UserId, &user.Nickname, &user.CreatTime)

	if err != nil {
		state = 0
		msg = "查询无数据"
	}else {
		state = 1
		msg = "查询完成"

		// 访问结构体并改变成员变量的值
		createTime := user.CreatTime
		createTime = common.DateToDate(createTime)
		user.CreatTime = createTime

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
		"content":   user,
	})
}
