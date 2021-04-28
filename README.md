#### Github：https://github.com/fyonecon/ginlaravel
#### Gitee：https://gitee.com/fyonecon/ginlaravel

## GinLaravel、GoLaravel
基于GO框架Gin开发，项目结构和项目哲学参考Laravel，支持MVC开发模式。本项目展示了MVC处理数据，同时展示"Safe—Controller—Kit"模式处理数据。

SCK模型大多数情况下是面向Api等，具有接口安全、请求数据直接易懂、复杂格式数据易操作（批量处理）、请求宽进严出等特点。SCK没有model层，也不推荐使用model。/Gen1/展示了使用MVC方式处理数据；/Gen3/展示了使用SCK处理数据和请求。

## 现已支持
Go-MySQL、Go-Redis、热更（fresh）、MVC、模版输出、Http访问频率拦截、HttpCors、对称加密（可中文）、http拦截器、Form上传文件。

## 理念
宽进严出，面向Api，能有效节约时间和保持秀发。

## 运行特点
内存常驻，可似swoole，Golang开发速度是PHP的十分之一，可支持热更。

## 项目目录解释
> -/app/Common/ 对go原生封装的公共函数、自定义必要数据参数配置

> -/app/Http/ 控制器、模型、拦截器

> -/app/Kit/ 自定义系统服务，包含第三方服务和系统功能服务

> -/bootstrap/ 系统服务启动、数据库全局引用配置，一般不需要更改此处。

> -/config/ MySQL数据库、Redis缓存、http端口号等配置

> -/routes/ 路由，同时支持模版型路由、Api型路由

> -/storage/ 系统日志、文件上传、静态缓存

> -/views/ 模版渲染的原始文件夹

> -ginlaravel 项目生产的二进制文件，在生产环境用

> -go.mod 项目所依赖的module路径、第三方库等的引入

> -server.go 跑起本项目的入口go-main文件。

## 命名原则
#### 自定义函数：大驼峰
#### 自定义变量：小驼峰
#### 自定义结构体和结构体成员：大驼峰
#### MySQL：小写+下划线
#### 接口名：小写+下划线
#### 自定义文件夹名：大驼峰
#### 系统集文件夹名：小写

## 运行项目
>Go运行环境。搭建Go和基础Gin环境请参考：https://blog.csdn.net/weixin_41827162/article/details/115693925

> MySQL（请将/项目资料/ginlaravel.sql 文件导入到数据库）

> cmd中运行"go run server.go"即可启动项目。或使用热更方式启动http服务，在cmd中目录运行"fresh"。

> 访问"http://127.0.0.1:8090/gen1/user/list_user "

> 项目上线：serverConfig["ENV"]的值改成release，然后使用以上同样方法运行。

## 如何初始化项目
以当前目录 /Users/fyonecon/go/src/ 为例
```sybase
获取源代码：
git clone https://github.com/fyonecon/ginlaravel.git
cd go/src/ginlaravel

使用 govendor 安装依赖包：
go get -u -v github.com/kardianos/govendor
govendor sync

初始化项目：
go mod init

构建依赖：
go mod tidy
go mod vendor

打包成二进制文件：
go build -mod=mod

启动MySQL数据库：请自行启动。
    
启动Redis：请自行启动。

在/config/mysql.go配置数据库信息，用来连接你的数据库。

启动http服务：
go run server.go


```

## 如何运行fresh热更服务
以项目目录 /Users/fyonecon/go/src/ginlaravel 为例
```sybase

去.bash_profile文件目录：
cd ~

运行：
source ~/.bash_profile

切换到项目目录：
cd go/src/ginlaravel

开启热更：
fresh

退出http服务用快捷键：Ctrl + C 。或直接关闭终端窗口。

```
以上即可项目开启的fresh热更服务。
若想一直开启终端窗口，请使用screen（yum install screen）来保持窗口。

热更服务文档：https://github.com/gravityblast/fresh 。

## （Mac环境）将项目打包成二进制文件
```sybase
项目根目录：
    go build -mod=mod

此时，在项目目录生成或更新了ginlaravel二进制文件。
开启文件的可执行权限：
    chmod 773 ginlaravel
        
在根目录运行：
    ./ginlaravel
        
即可开启二进制服务。

```

## 将项目部署在Centos7上（Go环境搭建、服务器环境搭建、数据库环境搭建）
教程：https://blog.csdn.net/weixin_41827162/article/details/116048754

## 关于
#### 作者Author：fyonecon
##### 长沙
##### 微信WeChat：fy66881159

## 版权
MIT