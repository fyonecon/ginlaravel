package Gen1App

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var initData = "initData=="

func Test2Run(ctx *gin.Context)  {
	fmt.Println("run")
	// 返回特殊格式意义的数据
	ctx.JSON(200, gin.H{
		"content": "run" + initData,
	})
}