package Runtime
// 硬件参数

// #include <unistd.h>
import "C"

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

func Hardware(ctx *gin.Context)  {

	memVirtual, _ := mem.VirtualMemory()
	fmt.Println("虚拟内存memVirtual=", memVirtual.Free/1024/1024)

	// 非常耗时
	//cpuPercent, _ := cpu.Percent(time.Second, false)
	//fmt.Println("cpuPercent=", cpuPercent)
	//println(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE), " bytes")

	//diskPart, _ := disk.Partitions(true)
	//fmt.Println("diskPart=", diskPart)

	ctx.Next()
}
