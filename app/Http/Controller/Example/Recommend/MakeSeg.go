package Recommend
// 批量生成分词库数据

import (
	"fmt"
	"ginvel.com/app/Common"
	"ginvel.com/app/Kit"
	KitLib2 "ginvel.com/app/Kit/KitLib"
	"github.com/gin-gonic/gin"
	"math"
)


type Keys struct { // 数据库表的字段（输出结果）
	NewsId int `json:"news_id"`
	Content  string `json:"content"`
}

func GoSave(theNewsId int)  {
	classId := "10"

	var res Keys

	// 多查询条件，仅限"="查询
	WhereMap := map[string]interface{}{}
	WhereMap["news_id"] = theNewsId
	Kit.DB.Table("gl_news").Where(WhereMap).First(&res)

	content := res.Content

	fmt.Println(theNewsId, classId)

	// 生成分词结果
	wordArray := KitLib2.MakeSeg(content)
	// 存储分词
	_ = SaveSegment(Common.IntToString(int64(theNewsId)), classId, wordArray)

}

// RunRoutine 多线程步长区块
func RunRoutine(f int, theStart int, theEnd int)  {
	fmt.Printf("~ 本次线程参数：线程index=%d，本次开始=%d，本次结束=%d \n", f, theStart, theEnd)
	for l:= theStart; l< theEnd; l++ {
		//time.Sleep(123 * time.Millisecond)
		fmt.Printf("@ 线程内循环：线程index=%d，内循环index=%d \n", f, l)
		// 其他操作
		GoSave(l)
	}
}

// SaveSeg1 文章分词及其分词+id入库
func SaveSeg1(ctx *gin.Context) {
	// GoSave(theNewsId)

	start := 3433 // 开始
	end := 4100 // 结束
	goNum := 40 // 多线程数量

	// 固定参数范围
	num := end-start
	if num < 0 {
		end = start
		goNum = 1
	}else if num < goNum {
		goNum = end-start
	}

	// 区间步长
	theFoot := math.Floor(float64((end-start) / goNum))
	// 处理主要
	for f:=0; f<goNum; f++ { // 多线程处理
		fmt.Printf("> 线程数量进度=%d/%d \n", f, goNum)
		theStart := float64(start) + theFoot*float64(f)
		theEnd := float64(start) + theFoot*float64(f+1)
		go RunRoutine(f, int(theStart), int(theEnd)) // 开启多线程
	}
	// 处理剩余
	theStart := start + int(theFoot)*goNum
	theEnd := end
	fmt.Printf("> 线程数量进度=%d完成，剩余数量=%d \n", goNum, theEnd - theStart)
	RunRoutine(goNum, theStart, theEnd)

	// 接口返回
	ctx.JSON(200, gin.H{
		"state": 1,
		"msg": "Sego中文分词示例",
		"content": map[string]interface{}{},
	})
}

