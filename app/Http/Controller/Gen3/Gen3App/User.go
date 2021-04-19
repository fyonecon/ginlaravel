package Gen3App

import (
	"database/sql"
	"ginlaravel/app/Common"
	"ginlaravel/app/Kit"
	"ginlaravel/bootstrap/driver"
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

	// 预定义参数
	var state int
	var msg string
	var testData map[string]string

	_page := Kit.Input(ctx, "page")
	_nickname := Kit.Input(ctx, "nickname")

	// 处理分页
	var limit int = Common.Page["limit"]
	var page int = Common.StringToInt(_page)
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
			newCreateTime := Common.DateToDate(theCreateTime)
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


