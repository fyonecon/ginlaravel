package Govern
// 硬件参数

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"
)

func GoGovern(ctx *gin.Context)  {

	memVirtual, _ := mem.VirtualMemory()
	fmt.Println("memVirtual=", memVirtual.Free/1024/1024)

	cpuPercent, _ := cpu.Percent(time.Second, false)
	fmt.Println("cpuPercent=", cpuPercent)

	//diskPart, _ := disk.Partitions(true)
	//fmt.Println("diskPart=", diskPart)

	ctx.Next()
}
