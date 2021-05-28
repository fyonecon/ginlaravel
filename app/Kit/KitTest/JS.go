package KitTest
// Go解释JavaScript

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func TestJS()  {
	vm := otto.New()
	//vm.Run()这里面是javascript代码,下面这段代码，是QQ空间提取出来的。
	vm.Run(`
        function u(x, K) {
        x += '';
        for (var N = [], T = 0; T < K.length; T++) N[T % 4] ^= K.charCodeAt(T);
        var U = ['EC', 'OK'],
        V = [];
        V[0] = x >> 24 & 255 ^ U[0].charCodeAt(0);
        V[1] = x >> 16 & 255 ^ U[0].charCodeAt(1);
        V[2] = x >> 8 & 255 ^ U[1].charCodeAt(0);
        V[3] = x & 255 ^ U[1].charCodeAt(1);
        U = [];
        for (T = 0; T < 8; T++) U[T] = T % 2 == 0 ? N[T >> 1] : V[T >> 1];
        N = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'];
        V = '';
        for (T = 0; T < U.length; T++) {
            V += N[U[T] >> 4 & 15];
            V += N[U[T] & 15]
        }
        return V
        }

        `)
	//使用vm.Call(函数名,nil,传递的参数,如果后面有多个参数，用逗号隔开就可以了)
	value, _ := vm.Call("u", nil, "12345678", "234")
	fmt.Println(value)
}
