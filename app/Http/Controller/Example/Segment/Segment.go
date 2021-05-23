package Segment

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	KitLib2 "ginvel.com/app/Kit/KitLib"
	"github.com/gin-gonic/gin"
	"github.com/go-ego/gse"
	"strings"
)

// SaveSeg1 文章分词及其分词+id入库
func SaveSeg1(ctx *gin.Context) {

	// 获取分词样本
	id := "4" // 对应样本的那个id
	text := []byte("汉服市场今年销售规模或将超过百亿")

	// 生成分词结果
	wordArray := MakeSeg(text)

	// 存储分词
	for h:=0; h<len(wordArray); h++ {
		txt := wordArray[h]

		type Keys struct { // 数据库表的字段（输出结果）
			SegDataId int `json:"seg_data_id"`
			Txt string `json:"txt"`
			Ids string `json:"ids"`
			UpdateTime string `json:"update_time"`
		}
		var res Keys

		fmt.Println(res)

		// 多查询条件，仅限"="查询
		WhereMap := map[string]interface{}{}
		WhereMap["txt"] = txt
		Kit.DB.Table("gl_seg_data").Where(WhereMap).First(&res)

		segDataId := res.SegDataId
		ids := res.Ids

		fmt.Println(res, txt, segDataId, ids)

		updateTime := Common.GetTimeDate("YmdHis")
		if len(ids) > 0 {
			_id := "-" + id + "-"
			if !strings.Contains(ids, _id) {
				// 更新数据
				newIds := ids + id + "-"
				// 新数据
				data := map[string]interface{}{
					"txt": txt,
					"ids": newIds,
					"update_time": updateTime,
				}
				_ = Kit.DB.Table("gl_seg_data").Where(WhereMap).Updates(data).RowsAffected
			}
		}else {
			// 新增数据
			newIds := "-" + id + "-"
			// 新数据
			data := Keys{
				Txt: txt,
				Ids: newIds,
				UpdateTime: updateTime,
			}

			_ = Kit.DB.Table("gl_seg_data").Select("txt", "ids", "update_time").Create(&data).Error
		}

	}

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "Sego中文分词示例",
		"content": map[string]interface{}{
			"date": wordArray,
			"id": id,
		},
	})
}

// FilterValueStringArray 删除数组中不能要的值
func FilterValueStringArray(arr []string, filterValue string) (newArr []string)  {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		theValue := arr[i]
		if theValue != filterValue && len(filterValue) > 0 {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// FilterLenStringArray 删除数组中不能要的值
func FilterLenStringArray(arr []string, _len int) (newArr []string)  {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		theValue := arr[i]
		if len(theValue) >= _len && _len > 0 {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// MakeSeg 生成文章的分词
func MakeSeg(text []byte) []string {

	// 分词方式1
	//var seg sego.Segmenter
	//seg.LoadDictionary(Common.ServerInfo["storage_path"] + "seg/data/dictionary.txt")
	//segments := seg.Segment(text)
	//wordSeg := sego.SegmentsToString(segments, false) // 支持普通模式和搜索模式两种分词，false普通模式

	// 分词方式2
	var seg gse.Segmenter
	seg.LoadDict(Common.ServerInfo["storage_path"] + "gse/data/dict/zh/dict.txt")
	seg.LoadStop()
	wordSeg := seg.String(string(text), false) // 支持普通模式和搜索模式两种分词，false普通模式

	// 过滤分词
	wordArray := strings.Split(wordSeg, " ")
	wordArray = Common.RemoveRepeatedStringArray(wordArray)
	wordArray = FilterLenStringArray(wordArray, 6)
	wordArray = FilterValueStringArray(wordArray, "的/uj")
	wordArray = FilterValueStringArray(wordArray, "得/ud")

	//wordArray = FilterValueStringArray(wordArray, " /x")
	//wordArray = FilterValueStringArray(wordArray, "./x")
	//wordArray = FilterValueStringArray(wordArray, ",/x")
	//wordArray = FilterValueStringArray(wordArray, ";/x")
	//wordArray = FilterValueStringArray(wordArray, "。/x")
	//wordArray = FilterValueStringArray(wordArray, "，/x")
	//wordArray = FilterValueStringArray(wordArray, "\n/x")
	//wordArray = FilterValueStringArray(wordArray, "、/x")
	//wordArray = FilterValueStringArray(wordArray, "?/x")
	//wordArray = FilterValueStringArray(wordArray, "？/x")
	//wordArray = FilterValueStringArray(wordArray, "!/x")
	//wordArray = FilterValueStringArray(wordArray, "！/x")
	//wordArray = FilterValueStringArray(wordArray, ":/x")
	//wordArray = FilterValueStringArray(wordArray, "：/x")

	return wordArray
}


// SearchSeg1 搜索分词结果
func SearchSeg1(ctx *gin.Context)  {
	// 查询条件
	text := "2021考研报名的步骤"
	// 生成分词结果
	wordArray := MakeSeg([]byte(text))

	// 分词id集合
	var allIds string
	for h:=0; h<len(wordArray); h++ {
		txt := wordArray[h]

		type Keys struct { // 数据库表的字段（输出结果）
			SegDataId int `json:"seg_data_id"`
			Txt string `json:"txt"`
			Ids string `json:"ids"`
			UpdateTime string `json:"update_time"`
		}
		var res Keys

		fmt.Println(res)

		// 多查询条件，仅限"="查询
		WhereMap := map[string]interface{}{}
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
	allIdsStringArray := strings.Split(allIds, "-")
	var allIdsIntArray []int
	for a:=0; a<len(allIdsStringArray); a++ {
		theValue := Common.StringToInt(allIdsStringArray[a])
		if theValue != 0 {
			allIdsIntArray = append(allIdsIntArray, int(theValue))
		}
	}

	rec := KitLib2.Recommend{}
	allIdsIntArray = rec.OrderHasArray(allIdsIntArray)

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "Sego中文分词搜索示例",
		"content": map[string]interface{}{
			"txt": wordArray,
			"id_array": allIdsIntArray,
		},
	})
}




// Seg2 中、英、日分词
func Seg2(ctx *gin.Context)  {
	var seg gse.Segmenter
	seg.LoadDict(Common.ServerInfo["storage_path"] + "gse/data/dict/zh/dict.txt")
	seg.LoadStop()
	// seg.LoadDictEmbed()
	// seg.LoadStopEmbed()

	text1 := "你好世界, Hello world"
	fmt.Println(seg.String(text1, false))

	segments := seg.Segment([]byte(text1))
	fmt.Println(gse.ToString(segments))
}
