## GinLaravel、GoLaravel
基于GO框架Gin开发，项目结构和项目哲学参考Laravel，支持MVC开发模式。本项目展示了MVC处理数据，同时展示"Safe——Controller——Kit"模式处理数据。

SCK模型大多数情况下是面向Api等，具有接口安全、请求数据直接易懂、复杂格式数据易操作（批量处理）、请求宽进严出等特点。SCK没有model层，也不推荐使用model。/Gen1Controller/展示了使用MVC方式处理数据；/Gen2Controller/展示了使用SCK处理数据和请求。

## 理念
宽进严出，面向Api，能有效节约时间和保持秀发。

## 运行特点
内存常驻，可似swoole，Golang开发速度是PHP的十分之一。

## 项目目录解释
> -/app/Common/ 对go原生封装的公共函数文件夹

> -/app/Http/ 控制器、模型、拦截器文件夹

> -/app/Kit/ 自定义系统服务，包含第三方服务和系统功能服务

> -/bootstrap/ 系统服务启动

> -/config/ 数据库配置、端口号配置等

> -/routes/ 路由，支持模版型路由、Api型路由

> -/views/ 模版渲染的原始文件夹

> -go.mod 项目所依赖的module路径、第三方库等的引入

> -server.go 跑起来本项目的入口go文件。

## 命名原则
#### 自定义函数：大驼峰
#### 自定义变量：小驼峰
#### 自定义结构体和结构体成员：大驼峰
#### MySQL：小写+下划线
#### 接口名：小写+下划线
#### 自定义文件夹名：大驼峰
#### 系统集文件夹名：小写

## 运行项目
>Go运行环境。搭建请参考：https://blog.csdn.net/weixin_41827162/article/details/115693925

> MySQL（请将/项目资料/ginlaravel.sql 文件导入到数据库）

> cmd中运行"go run server.go"即可启动项目。或使用热更方式启动http服务，在cmd中目录运行"fresh"。

> 访问"http://127.0.0.1:8090/gen2/app/list_user "

> 项目上线：serverConfig["ENV"]的值改成release，然后使用以上同样方法运行。

## 关于
#### 第二作者Author：fyonecon
#### 博客Blog：https://blog.csdn.net/weixin_41827162/article/details/115712700
#### Github：https://github.com/fyonecon/ginlaravel
##### 长沙
##### 微信WeChat：fy66881159

## 版权
MIT