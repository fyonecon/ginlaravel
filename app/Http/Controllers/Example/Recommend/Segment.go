package Recommend

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	"ginvel.com/app/Kit/KitLib"
	"strings"
)

// SaveSegment 保存分词
func SaveSegment(id string, classId string, wordArray []string) int {
	fmt.Println("id=" + id)
	_id := Common.StringToInt(id)
	if _id == 0 {
		fmt.Println("id格式不正确=" + id)
		return 0
	}
	// 存储分词
	for h:=0; h<len(wordArray); h++ {
		txt := wordArray[h]

		type Keys struct { // 数据库表的字段（输出结果）
			SegDataId int `json:"seg_data_id"`
			ClassId  string `json:"class_id"`
			Txt string `json:"txt"`
			Ids string `json:"ids"`
			UpdateTime string `json:"update_time"`
		}
		var res Keys

		fmt.Println(res)

		// 多查询条件，仅限"="查询
		WhereMap := map[string]interface{}{}
		WhereMap["class_id"] = classId
		WhereMap["txt"] = txt
		Kit.DB.Table("gl_seg_data").Where(WhereMap).First(&res)

		segDataId := res.SegDataId
		ids := res.Ids

		fmt.Println(res, txt, segDataId, ids)

		updateTime := Common.GetTimeDate("YmdHis")
		if len(ids) > 0 {
			_id := "_" + id + "_"
			if !strings.Contains(ids, _id) {
				// 更新数据
				newIds := ids + id + "_"
				// 新数据
				data := map[string]interface{}{
					"class_id":    classId,
					"txt":         txt,
					"ids":         newIds,
					"update_time": updateTime,
				}
				_ = Kit.DB.Table("gl_seg_data").Where(WhereMap).Updates(data).RowsAffected
			}
		}else {
			// 新增数据
			newIds := "_" + id + "_"
			// 新数据
			data := Keys{
				ClassId: classId,
				Txt: txt,
				Ids: newIds,
				UpdateTime: updateTime,
			}

			_ = Kit.DB.Table("gl_seg_data").Select("class_id", "txt", "ids", "update_time").Create(&data).Error
		}

	}

	return 1
}

// SearchSegment 根据分词查询目标权重id集合
func SearchSegment(classId string, wordArray []string) []int {
	var allIds string
	for h:=0; h<len(wordArray); h++ {
		txt := wordArray[h]

		type Keys struct { // 数据库表的字段（输出结果）
			SegDataId int `json:"seg_data_id"`
			ClassId  string `json:"class_id"`
			Txt string `json:"txt"`
			Ids string `json:"ids"`
			UpdateTime string `json:"update_time"`
		}
		var res Keys

		fmt.Println(res)

		// 多查询条件，仅限"="查询
		WhereMap := map[string]interface{}{}
		WhereMap["class_id"] = classId
		WhereMap["txt"] = txt
		Kit.DB.Table("gl_seg_data").Where(WhereMap).First(&res)

		segDataId := res.SegDataId
		ids := res.Ids

		fmt.Println(res, txt, segDataId, ids)

		if len(ids) > 0 {
			// 有分词结果，则进行分词查询
			allIds = allIds + ids
		}

	}

	// 处理分词集合
	allIdsStringArray := strings.Split(allIds, "_")
	var allIdsIntArray []int
	for a:=0; a<len(allIdsStringArray); a++ {
		theValue := Common.StringToInt(allIdsStringArray[a])
		if theValue != 0 {
			allIdsIntArray = append(allIdsIntArray, int(theValue))
		}
	}

	rec := KitLib.Recommend{}
	allIdsIntArray = rec.OrderHasArray(allIdsIntArray)

	return allIdsIntArray
}
