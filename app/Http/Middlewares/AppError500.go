package Middlewares
// 接管服务器500错误，使错误可视化

import (
	"fmt"
	"ginvel.com/app/Common"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

// HideAppInfo 隐藏必要关键词
func HideAppInfo( _info string) (info string)  {
	_info = strings.Replace(_info, "ginvel.com", "*", -1)
	_info = strings.Replace(_info, "laravel", "*", -1)
	_info = strings.Replace(_info, "go", "*", -1)
	_info = strings.Replace(_info, "src", "*", -1)
	_info = strings.Replace(_info, "pkg", "*", -1)
	_info = strings.Replace(_info, "app", "**", -1)
	_info = strings.Replace(_info, "Http", "**", -1)
	_info = strings.Replace(_info, "/", "*", -1)
	_info = strings.Replace(_info, ".", "*", -1)
	_info = strings.Replace(_info, "Middleware", "*", -1)
	_info = strings.Replace(_info, "Ruler", "*", -1)
	_info = strings.Replace(_info, "Common", "*", -1)
	_info = strings.Replace(_info, "Runtime", "*", -1)
	_info = strings.Replace(_info, "Task", "*", -1)
	_info = strings.Replace(_info, "Controller", "*", -1)
	_info = strings.Replace(_info, "Model", "*", -1)
	_info = strings.Replace(_info, "Kit", "*", -1)
	_info = strings.Replace(_info, "extent", "*", -1)
	_info = strings.Replace(_info, "config", "*", -1)
	_info = strings.Replace(_info, "bootstrap", "*", -1)

	info = _info
	return
}

// AppError500 抛出500错误
func AppError500 (ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			fmt.Printf("panic: %v\n", err)
			// debug.PrintStack() // 显示报错详情

			pc := make([]uintptr, 8)
			runtime.Callers(2, pc)
			//f := runtime.FuncForPC(pc[0])

			_fc0 := runtime.FuncForPC(pc[0]).Name()
			_fc1 := runtime.FuncForPC(pc[1]).Name()
			_fc2 := runtime.FuncForPC(pc[2]).Name()
			_fc3 := runtime.FuncForPC(pc[3]).Name()
			_fc4 := runtime.FuncForPC(pc[4]).Name()
			_fc5 := runtime.FuncForPC(pc[5]).Name()
			//_fc6 := runtime.FuncForPC(pc[6]).Name()

			fc0 := HideAppInfo(_fc0)
			fc1 := HideAppInfo(_fc1)
			fc2 := HideAppInfo(_fc2)
			fc3 := HideAppInfo(_fc3)
			fc4 := HideAppInfo(_fc4)
			fc5 := HideAppInfo(_fc5)
			//fc6 = HideAppInfo(_fc6)

			errorFunc1 := gin.H{
				"0": _fc0,
				"1": _fc1,
				"2": _fc2,
				"3": _fc3,
				"4": _fc4,
				"5": _fc5,
				//"6": _fc6,
			}
			errorFunc2 := gin.H{
				"0": fc0,
				"1": fc1,
				"2": fc2,
				"3": fc3,
				"4": fc4,
				"5": fc5,
				//"6": fc6,
			}

			fmt.Println(errorFunc1)

			// 返回
			ctx.JSON(500, gin.H{
				"state": 500,
				"msg": "代码运行报错，请查看代码运行日志",
				"content": gin.H{
					"gl_version": Common.ServerInfo["gl_version"],
					"error_func": errorFunc2,
					"error_msg": err,
				},
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用并返回JSON提示
	ctx.Next()
}


