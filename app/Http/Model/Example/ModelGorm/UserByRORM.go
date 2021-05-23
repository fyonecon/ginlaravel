package ModelGorm

import (
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
)

// ListUserKeys 查用户列表
type ListUserKeys struct { // 参数名需大写
	UserId int64 `json:"user_id"`
	Nickname string `json:"nickname"`
	UserClassId int64 `json:"user_class_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
func (model *ListUserKeys) ListUser(_limit int, _offset int, _userClassId int64, _nickname string) (res []ListUserKeys, total int64, err error) {

	// 多查询条件，仅限"="查询
	WhereMap := map[string]interface{}{}
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2},
	}

	// 数据列表
	list := Kit.DB.Table("gl_user").Not(NotMap).Where(WhereMap)
	// 数据总数
	count := Kit.DB.Table("gl_user").Not(NotMap).Where(WhereMap)

	// 加入非"="多条件查询
	if len(_nickname) != 0 {
		list = list.Where("nickname LIKE ?", "%" + _nickname + "%")
		count = count.Where("nickname LIKE ?", "%" + _nickname + "%")
	}
	if _userClassId != 0 {
		list = list.Where("user_class_id = ?", _userClassId)
		count = count.Where("user_class_id = ?", _userClassId)
	}

	// 完成其他条件
	list.Order("nickname asc").Limit(_limit).Offset(_offset).Scan(&res)
	count.Order("nickname asc").Count(&total)

	// 遍历struct数据
	for i := 0; i < len(res); i++ {
		theCreateTime := res[i].CreateTime
		newCreateTime := Common.DateToDate(theCreateTime)
		res[i].CreateTime = newCreateTime

	}

	return
}

// ThatUserKeys 某用户
type ThatUserKeys struct { // 数据库表的字段（输出结果）
	UserId int64 `json:"user_id"`
	Nickname string `json:"nickname"`
	UserClassId int64 `json:"user_class_id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
func (model *ThatUserKeys) ThatUser(userId int64) (res ThatUserKeys, err error)  {

	// 多查询条件，仅限"="查询
	WhereMap := map[string]interface{}{}
	WhereMap["user_id"] = userId
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2},
	}

	// 操作数据库
	Kit.DB.Table("gl_user").Not(NotMap).Where(WhereMap).First(&res)

	createTime := res.CreateTime
	createTime = Common.DateToDate(createTime)
	res.CreateTime = createTime

	return
}

// AddUserKeys 新增
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

	err = Kit.DB.Table("gl_user").Select("user_class_id", "nickname", "create_time").Create(&data).Error

	_id = data.ID

	return
}

// UpdateUserKeys 更新
type UpdateUserKeys struct {}
func (model *UpdateUserKeys)UpdateUser(userId int64, userClassId int64, nickname string, updateTime string) (res int64, err error) {

	// 多查询条件，仅限"="查询
	WhereMap := map[string]interface{}{}
	WhereMap["user_id"] = userId
	// 多排除查询条件
	NotMap := map[string]interface{}{
		"state": []int64{2}, // state <> [2, 3, 4]
	}

	// 新数据
	data := map[string]interface{}{
		"user_class_id": userClassId,
		"nickname": nickname,
		"update_time": updateTime,
	}

	res = Kit.DB.Table("gl_user").Not(NotMap).Where(WhereMap).Updates(data).RowsAffected

	return
}

// DelUserKeys 删除用户。不是真正删除，仅不可见状态
type DelUserKeys struct {}
func (model *DelUserKeys)DelUser(userId int64, updateTime string) (res int64, err error) {

	// 多查询条件，仅限"="查询
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

	res = Kit.DB.Table("gl_user").Not(NotMap).Where(WhereMap).Updates(data).RowsAffected

	return
}

// ClearUserKeys 彻底删除用户，按条件删除
type ClearUserKeys struct {}
func (model *ClearUserKeys)ClearUser(userId int64,) (res int64, err error) {

	// 多查询条件，仅限"="查询
	WhereMap := map[string]interface{}{}
	//WhereMap["state"] = 1
	WhereMap["user_id"] = userId

	// 新数据
	data := ClearUserKeys{}

	res = Kit.DB.Table("gl_user").Where(WhereMap).Delete(data).RowsAffected

	return
}

