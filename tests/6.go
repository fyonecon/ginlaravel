package main
// 链式操作
// https://learnku.com/docs/the-way-to-go/1412-chaining-goroutines/3696
import (
	"flag"
	"fmt"
)

var numGoroutine = flag.Int("num", 99, "how many goroutines")

func Channal(left, right chan int) {
	left <- 1 + <-right
}

func main() {

	flag.Parse()

	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *numGoroutine; i++ {
		left, right = right, make(chan int)
		go Channal(left, right) //
	}

	right <- 0

	// start the chaining 开始链接
	x := <-leftmost // wait for completion 等待完成
	fmt.Println(x)

}