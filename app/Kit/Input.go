package Kit

import (
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
	"strings"
)

/ Input 自动判断请求类型并自动输出参数值
// 针对REST风格的GET和普通的POST
/*

// 方式1：fetch
// POST请求
const post_api = api_url + "admin/admin_login"; // 接口
const map = new Map([ // 要提交数据
  ["app_class", "test"],
  ["url", encodeURIComponent(window.location.href).substring(0, 2000)], // 取当前url即可
]);
let body = "";
for (let [k, v] of map) { body += k+"="+v+"&"; } // 拼装数据，限制2MB
fetch(post_api, {
  method: "post",
  mode: "cors", // same-origin/no-cors/cors
  cache: "no-cache",
  headers: {
	  "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
  },
  body: body,
}).then(function(response){
  if (response.status === 200){return response;}
}).then(function(data) {
  return data.text();
}).then(function(text){ // 返回接口数据
  //

}).catch(function(error){
  let error_info = "Fetch_Error：" + error;
});
// 结束-Fetch

// axios拼装POST请求

// 方式2：axios
// POST请求
const post_api = ""; // api
const map = new Map([ // 要提交数据
	["app_class", "test"],
	["url", encodeURIComponent(window.location.href).substring(0, 2000)],
]);
const body = new URLSearchParams();
for (let [k, v] of map) { body.append(k, v+""); }
axios.post(post_api, body, {headers: {"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"}})
	.then(function (back) {
		let res = back.data;
		console.log(res)
	})
	.catch(function (e) {
		console.error(e);
	});

// 方式3：axios
// POST请求
axios.post(
  api_url + "admin/admin_login", // 设置了baseUrl就不需要连接主域名
  {
	app_class: "test",
	url: encodeURIComponent(window.location.href).substring(0, 2000),
  },
  {
	headers: {"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"},
  }
)
.then(function (back) {
  let res = back.data;

})
.catch(function (e) {
  console.error(e);
});


// 注意：「application/json」下需要前端自己调跨域，但是还是不建议用此。推荐使用「application/x-www-form-urlencoded 」或 「 multipart/form-data 」。


*/
func Input(ctx *gin.Context, key string) string {
	var value string
	var hasKey bool

	_method := ctx.Request.Method
	_contentType := ctx.Request.Header["Content-Type"]

	if _method == "GET" {
		value, hasKey = ctx.GetQuery(key)
	}else if _method == "POST" {
		if len(_contentType) >= 1 { // 判断是否含有请求头信息数组
			hasCt1 := strings.Contains(_contentType[0], "application/x-www-form-urlencoded") // 一般用于参数
			hasCt2 := strings.Contains(_contentType[0], "multipart/form-data") // 一般用于文件或参数
			if hasCt1{
				value, hasKey = ctx.GetPostForm(key)
			}else if hasCt2 {
				value, hasKey = ctx.GetPostForm(key)
			}else {
				helper.Log("POST方式时建议：Content-Type=「application/x-www-form-urlencoded 」或 「 multipart/form-data 」")
			}
		}else {
			helper.Log(_contentType)
		}
	}else if _method == "OPTIONS" {
		value, hasKey = "", false
	} else {
		value, hasKey = "", false
	}

	if hasKey == false { // 参数不存在就重新验证参数

		//
		formJson := ctx.Request.Form
		for formData, _ := range formJson{
			value = formData
			break
		}
		var v interface{}
		err := json.Unmarshal([]byte(value), &v)
		if err != nil {
			return ""
		}else {
			data := v.(map[string]interface{})
			value = helper.InterfaceToString(data[key])
		}

	}

	return value
}
