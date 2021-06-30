package main
// 使用goroutines开指定数量的线程池，通道分别传递任务和任务结果。
// https://www.cnblogs.com/swarmbees/p/6601145.html

import (
	"fmt"
	"runtime"
)

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
// jobs <-chan int：只能接收数据
// results chan<- int：只能发送数据
func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Println(id*i)
		fmt.Println("worker-w=", id, "started-job-i=", i)
		fmt.Println("worker-w=", id, "finished-job-i=", i)
		results <- i+1
	}
}

func main() {

	// 利用CPU多核
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int) // 只用于接收数据
	results := make(chan int) // 只用于发送数据
	loop := 100 // 线程数量

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	// 用于处理任务，并将任务通道和结果通道传递给了线程函数
	for g := 1; g <= loop; g++ {
		fmt.Println("g=", g)
		go worker(g, jobs, results)
	}
	fmt.Println("go - 循环完成")

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	// 发送任务到jobs通道，工作线程在没有任务时，是阻塞着等待任务。
	// 当发现任务通道中有任务时，开始执行任务，当任务执行完毕时，将任务结果发送给结果通道。
	for r := 1; r <= loop; r++ {
		fmt.Println("r=", r)
		jobs <- r
	}
	close(jobs)
	fmt.Println("jobs - 循环完成")

	// Finally we collect all the results of the work.
	for w := 1; w <= loop; w++ {
		fmt.Println("w=", w)
		<-results
	}
	fmt.Println("results - 循环完成")
}

