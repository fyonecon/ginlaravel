package Recommend

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	KitLib2 "ginvel.com/app/Kit/KitLib"
	"github.com/gin-gonic/gin"
)

// SearchSeg1 搜索分词结果
func SearchSeg1(ctx *gin.Context)  {

	// 处理分页
	_page := Kit.Input(ctx, "page")
	page := int(Common.StringToInt(_page))
	var limit int = Common.Page["limit"]
	var offset int = 0 // 本页从第几个开始
	if page <= 0 { page = 1 } else if page > 200 { page = 200 }
	page = page - 1
	offset = limit*page

	// 查询条件
	classId := "10"
	text := Kit.Input(ctx, "text")
	if len(text) == 0 {
		text = "课程"
	}

	// 生成分词结果
	wordArray := KitLib2.MakeSeg(text)
	//wordArray := []string{"考研/vn","住宿/n","环境/n","收费/n"}
	fmt.Println(wordArray)
	// 分词id集合
	allIdsIntArray := SearchSegment(classId, wordArray)

	// 利用分页切割id数组
	total := len(allIdsIntArray)
	paging := Common.MakePaging(int(total), limit, page)
	idArray := allIdsIntArray[offset: offset+limit] // 本页内的词条id

	// 获取运行耗时，ms
	StatStart, _ := ctx.Get("stat_start")
	StatLatency, _ := ctx.Get("stat_latency")

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "Sego中文分词搜索示例",
		"paging": paging,
		"content": map[string]interface{}{
			"txt": wordArray,
			"id": idArray,
			"state_start": StatStart,
			"state_latency": StatLatency,
			//"id_array": allIdsIntArray,
		},
	})
}



