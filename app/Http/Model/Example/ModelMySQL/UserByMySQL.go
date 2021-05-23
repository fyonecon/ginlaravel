package ModelMySQL

import (
	"database/sql"
	"ginvel.com/app/Common"
	"ginvel.com/bootstrap/driver"
	"log"
)

var db *sql.DB = driver.MysqlDb

// ListUserKeys 查用户列表
type ListUserKeys struct { // 参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func (model *ListUserKeys) ListUser(_page int, _userClassId int, _nickname string) (users []ListUserKeys, err error) {

	// 处理分页
	var limit int = Common.Page["limit"]
	var page int = 0
	var offset int = 0 // 本页从第几个开始
	if _page <= 0 { page = 1 } else if _page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 构建查询
	nickname := "%" + _nickname + "%" // 模糊查询

	// 查询数据
	users = make([]ListUserKeys, 0)
	rows, err := db.Query("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 and `nickname` like ? LIMIT ?, ?", nickname, offset, limit)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	var user ListUserKeys
	for rows.Next() {
		rows.Scan(&user.UserId, &user.Nickname, &user.CreatTime)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}

	return
}

// ThatUserKeys 查某用户
type ThatUserKeys struct { // 参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func (model *ThatUserKeys) ThatUser(userId int) (user ThatUserKeys, err error)  {
	err = db.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 and `user_id`=?", userId).Scan(&user.UserId, &user.Nickname, &user.CreatTime)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}

// AddUserKeys 新增
type AddUserKeys struct {}
func (model *AddUserKeys)AddUser(userClassId int64, nickname string, createTime string) (_id int64, err error)  {
	data, err := db.Exec("INSERT INTO `gl_user`(`user_class_id`, `nickname`, `create_time`) VALUES (?, ?, ?)", userClassId, nickname, createTime)

	if err != nil {
		log.Println(err.Error())
		return
	}

	_id, err = data.LastInsertId()
	return
}

// UpdateUserKeys 更新
type UpdateUserKeys struct {}
func (model *UpdateUserKeys)UpdateUser(userId int64, userClassId int64, nickname string, updateTime string) (res int64, err error) {
	data, err := db.Exec("UPDATE `gl_user` SET `user_class_id`=?, `nickname`=?, `update_time`=? WHERE `state`=1 and `user_id` = ?", userClassId, nickname, updateTime, userId)

	if err != nil {
		log.Println(err.Error())
		return
	}

	res, err = data.RowsAffected()
	return
}

// DelUserKeys 删除用户
// 不是真正删除，仅不可见状态
type DelUserKeys struct {}
func (model *DelUserKeys)DelUser(userId int64, updateTime string) (res int64, err error) {
	data, err := db.Exec("UPDATE `gl_user` SET `state`=2, `update_time`=? WHERE `state`=1 and `user_id` = ?", updateTime, userId)

	if err != nil {
		log.Println(err.Error())
		return
	}

	res, err = data.RowsAffected()
	return
}

// ClearUserKeys 彻底删除用户
type ClearUserKeys struct {}
func (model *ClearUserKeys)ClearUser(userId int64,) (res int64, err error) {
	data, err := db.Exec("DELETE FROM `gl_user` WHERE `user_id` = ?", userId)

	if err != nil {
		log.Println(err.Error())
		return
	}

	res, err = data.RowsAffected()
	return
}


