package main

import (
	"fmt"
)

func main()  {

	var str string = "2021"
	var div string = `
<div style="color:blue;" data-id='1'>
	<script>alert("js")</script>
	<!-- 动态渲染html内容 -->
	<h3 class="red">
		` + str + `
	</h3>
</div>
	`

	fmt.Println(div)

}
