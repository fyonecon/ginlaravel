package Gen3

import (
	"ginlaravel/bootstrap/driver"
	"gorm.io/gorm"
)

var gDB *gorm.DB = driver.GDB // 连接gorm扩展

// 某用户
type ThatUserKeys struct { // 数据库表的字段（输出结果）
	UserId int `json:"user_id"`
	Nickname string `json:"nickname"`
	UserClassId string `json:"user_class_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
func (model *ThatUserKeys) ThatGMUser(userId int64) (res ThatUserKeys, err error)  {

	// 多查询条件
	WhereMap := map[string]interface{}{}
	WhereMap["user_id"] = userId
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2},
	}

	// 操作数据库
	res = ThatUserKeys{}
	gDB.Table("gl_user").Not(NotMap).Where(WhereMap).Find(&res)

	return
}

// 新增
type AddUserKeys struct {
	ID int64 // 要返回的主键值
	UserClassId int64 `json:"user_class_id"`
	Nickname    string `json:"nickname"`
	CreateTime  string `json:"create_time"`

}
func (model *AddUserKeys)AddUser(userClassId int64, nickname string, createTime string) (_id int64, err error)  {

	// 新数据
	data := AddUserKeys{
		UserClassId: userClassId,
		Nickname: nickname,
		CreateTime: createTime,
	}

	err = gDB.Table("gl_user").Select("user_class_id", "nickname", "create_time").Create(&data).Error

	_id = data.ID

	return
}

// 更新
type UpdateUserKeys struct {}
func (model *UpdateUserKeys)UpdateUser(userId int64, userClassId int64, nickname string, updateTime string) (res int64, err error) {

	// 多查询条件
	WhereMap := map[string]interface{}{}
	WhereMap["user_id"] = userId
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2, 3, 4}, // state <> [2, 3, 4]
	}

	// 新数据
	data := map[string]interface{}{
		"user_class_id": userClassId,
		"nickname": nickname,
		"update_time": updateTime,
	}

	res = gDB.Table("gl_user").Not(NotMap).Where(WhereMap).Updates(data).RowsAffected

	return
}

// 删除用户
// 不是真正删除，仅不可见状态
type DelUserKeys struct {}
func (model *DelUserKeys)DelUser(userId int64, updateTime string) (res int64, err error) {

	// 多查询条件
	WhereMap := map[string]interface{}{}
	WhereMap["user_id"] = userId
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2}, // state <> [2]
	}

	// 新数据
	data := map[string]interface{}{
		"state": 2,
		"update_time": updateTime,
	}

	res = gDB.Table("gl_user").Not(NotMap).Where(WhereMap).Updates(data).RowsAffected

	return
}

// 彻底删除用户，按条件删除
type ClearUserKeys struct {}
func (model *ClearUserKeys)ClearUser(userId int64,) (res int64, err error) {

	// 多查询条件
	WhereMap := map[string]interface{}{}
	//WhereMap["state"] = 1
	WhereMap["user_id"] = userId

	// 新数据
	data := ClearUserKeys{}

	res = gDB.Table("gl_user").Where(WhereMap).Delete(data).RowsAffected

	return
}
