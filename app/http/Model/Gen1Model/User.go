package Gen1Model

import (
	"database/sql"
	"ginlaravel/app/provider/driver"
	"log"
)

var db *sql.DB = driver.MysqlDb




// 用户列表
//type ListUserModel struct { // 参数名需大写
//	UserId int
//	Nickname string
//	CreatTime string
//}
//func (model *UserModel) ListUser(userClassId int, nickname string) (user UserModel, err error) {
//	err = db.QueryRow()
//
//
//}


// 哪个用户
type ThatUserModel struct { // 参数名需大写
	UserId int
	Nickname string
	CreatTime string
}
func (model *ThatUserModel) ThatUser(userId int) (user ThatUserModel, err error)  {
	 err = db.QueryRow("SELECT `user_id`, `nickname`, `create_time` FROM `gl_user` WHERE `state`=1 and `user_id` = ?", userId).Scan(&user.UserId, &user.Nickname, &user.CreatTime)
	if err != nil {
		log.Println(err.Error())
		return
	}
	return
}