package Gen1Model

import (
	"database/sql"
	"ginlaravel/app/common"
	"ginlaravel/app/provider/driver"
	"log"
)

var db *sql.DB = driver.MysqlDb

// 用户列表
type ListUserModel struct { // 参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func (model *ListUserModel) ListUser(_page int, _userClassId int, _nickname string) (users []ListUserModel, err error) {

	// 处理分页
	var limit int = common.Page["limit"]
	var page int = 0
	var offset int = 0 // 本页从第几个开始
	if _page <= 0 { page = 1 } else if _page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 构建查询
	//userClassId := ""
	//if _userClassId != 0 {
	//	__userClassId := common.IntToString(_userClassId)
	//	userClassId = " `user_class_id`=" + __userClassId
	//}else {
	//	userClassId = ""
	//}
	nickname := "%" + _nickname + "%"

	// 查询数据
	users = make([]ListUserModel, 0)
	rows, err := db.Query("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 and `nickname` like ? LIMIT ?, ?", nickname, offset, limit)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	var user ListUserModel
	for rows.Next() {
		rows.Scan(&user.UserId, &user.Nickname, &user.CreatTime)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}

	return
}


// 哪个用户
type ThatUserModel struct { // 参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func (model *ThatUserModel) ThatUser(userId int) (user ThatUserModel, err error)  {
	 err = db.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 and `user_id`=?", userId).Scan(&user.UserId, &user.Nickname, &user.CreatTime)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}