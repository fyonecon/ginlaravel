package Recommend

import (
	"fmt"
	"ginlaravel/app/Kit"
	"ginlaravel/extend/KitLib"
	"github.com/gin-gonic/gin"
)

// SearchSeg1 搜索分词结果
func SearchSeg1(ctx *gin.Context)  {
	// 查询条件
	classId := "10"
	text := Kit.Input(ctx, "text")
	if len(text) == 0 {
		text = "课程"
	}

	// 生成分词结果
	wordArray := KitLib.MakeSeg(text)
	//wordArray := []string{"考研/vn","住宿/n","环境/n","收费/n"}
	fmt.Println(wordArray)
	// 分词id集合
	allIdsIntArray := SearchSegment(classId, wordArray)

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



