package Task
// 硬件参数

// #include <unistd.h>
import "C"

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/shirou/gopsutil/v3/cpu"
	_ "github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load" // CPU负载
	"github.com/shirou/gopsutil/v3/mem"  // 内存占用
	_ "github.com/shirou/gopsutil/v3/net"
	"runtime"
	"time"
)

func Hardware()  {

	// 内存信息
	// total // 内存大小
	// available // 闲置可用内存
	// used // 已使用内存
	// usedPercent // 已使用百分比
	memVirtual, _ := mem.VirtualMemory()
	fmt.Println("虚拟内存free=", memVirtual.Free/1024/1024)
	fmt.Println("虚拟内存used=", memVirtual.Used/1024/1024)
	fmt.Println("虚拟内存UsedPercent=", memVirtual.UsedPercent/1024/1024)

	// 逻辑CPU数量
	cpuNum := runtime.NumCPU()
	fmt.Println("逻辑CPU数量=", cpuNum)
	Common.SetGlobalData("cpu_num", int64(cpuNum))

	// CPU使用率（此插件此功能比较耗时）
	cpuPercent, _ := cpu.Percent(time.Second, false)
	fmt.Println("cpuPercent=", cpuPercent)
	Common.SetGlobalData("cpu_percent", cpuPercent[0])

	// C语言
	println(C.sysconf(C._SC_PHYS_PAGES)*C.sysconf(C._SC_PAGE_SIZE), " bytes")

	// CPU负载（不耗时）
	cpuLoad, _ := load.Avg()
	// {"load1":3.62109375,"load5":2.93408203125,"load15":2.58251953125}
	// load表示每1分钟、5分钟、15分钟的平均队列（平均负载）,值为进程或线程数
	// 具体示意请参考load average：https://blog.csdn.net/bd_zengxinxin/article/details/51781630
	fmt.Printf("CPU负载：%v ", cpuLoad)
	Common.SetGlobalData("cpu_load1", cpuLoad.Load1)
	Common.SetGlobalData("cpu_load5", cpuLoad.Load1)
	Common.SetGlobalData("cpu_load15", cpuLoad.Load1)

	//diskPart, _ := disk.Partitions(true)
	//fmt.Println("diskPart=", diskPart)

	// net IO
	//io, _ := net.IOCounters(true)
	//for index, v := range io {
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	//}
}
