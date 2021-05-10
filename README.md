## 欢迎 Star🌟 本项目
#### Github（推荐）：https://github.com/fyonecon/ginlaravel
#### Gitee（备用）：https://gitee.com/fyonecon/ginlaravel

## 版本
#### v1.4

[comment]: <> (## 官网)

[comment]: <> (#### https://ginlaravel.com)

## GinLaravel（GoLaravel）介绍
基于Golang框架Gin（Gin学习文档：https://learnku.com/docs/gin-gonic/2019 ）开发，项目结构和项目理念参考Laravel。学习本项目时建议从阅读本项目源码开始，并具备Golang、Gin、go mod、centos+nginx、redis、mysql、加密解密、http、Swagger等必要知识。GinLaravel的构建和运行周期与Beego、Vue、React、Laravel、ThinkPHP、Django等都会有类似的引导思路、参数设置、插件扩展、服务部署、代码统一性、生态护城河等。GinLaravel尽量减少汉字造词、功能造词、贤者造词，不会用描述华丽、中英文混杂的造词方式去包装一个功能平常的老东西（比如 http+header、cookie/localstorage/sqlite+token、加密与解密、数据缓存+队列、负载均衡+docker、内存常驻与微服务、接口安全与前后端分离等）。

整个开发以人为本，功能模块化扩展，系统安全可控，渐进式提升访问量。

GinLaravel支持MVC开发模式。本项目展示了MVC处理数据，同时展示"Verify—Controller—Kit"模式处理数据。

## 现已支持
Go-MySQL、GORM（v2）、Go-Redis、热更（fresh）、Swagger、MVC、模版输出、Http访问频率拦截、HttpCors、对称加密（可中文）、http拦截器、多层路由等。

## 理念
宽进严出，面向Api，适合复杂项目，任何参数或服务都会有默认值。

## 运行特点
内存常驻似swoole，Golang开发速度是PHP的十分之一，支持fresh热更，空间命名分布明确适合复杂业务项目，整个项目运行需go+mysql+redis，遇到参数或服务不可用时不会立即中断请求。

## 项目目录解释
--/app/Common/ 对go原生封装的公共函数、自定义必要数据参数配置

--/app/Http/ 控制器、模型、拦截器

--/app/Http/Controller/ 控制器

--/app/Http/Controller/Example 示例

--/app/Http/Controller/Gen3 版本3的控制器文件夹

--/app/Http/Model/ 数据模型

--/app/Middleware/ 中间件，含有cors、http限速

--/app/Kit/ 自定义系统服务，包含第三方服务和系统功能服务

--/bootstrap/ 系统服务启动、数据库全局引用配置，一般不需要更改此处。

--/bootstrap/app/ 项目运行入口

--/bootstrap/driver/ 服务驱动

--/config/ MySQL数据库、Redis缓存、http端口号等配置

--/config/docs/ 接口文档插件Swagger的目录

--/routes/ 路由，同时支持模版型路由（web.go）、Api型路由（api.go）。推荐4层路由命名。

--/tmp/ fresh热更的缓存日志目录

--/storage/ 系统日志、文件上传、静态缓存。目录需权限777。

--/views/ 模版渲染的原始文件夹

--/views/pages/ 模版的html文件

--/views/static/ 模板静态资源：js、css、img

--ginlaravel 项目生产的二进制文件，在生产环境用

--go.mod 项目所依赖的module路径、第三方库等的引入

--server.go 跑起本项目的入口main文件。

## 命名原则
#### 自定义函数：大驼峰
#### 自定义变量：小驼峰
#### 自定义结构体和结构体成员：大驼峰
#### MySQL、Redis：小写+下划线
#### 接口名：小写+下划线
#### 自定义文件夹名：大驼峰
#### 系统集文件夹名：小写

## 运行项目
> Go运行环境。搭建Go和基础Gin环境请参考：https://blog.csdn.net/weixin_41827162/article/details/115693925

> MySQL（请将/项目资料/ginlaravel.sql 文件导入到数据库）

> cmd中运行"go run main.go"即可启动项目。或使用热更方式启动http服务，在cmd中目录运行"fresh"。v1.4开始为适配swaggo遂将server.go更名为main.go。

> 访问：http://127.0.0.1:8090

> 项目上线：serverConfig["ENV"]的值改成release，然后使用以上同样方法运行。

> 现已测试过的客户端环境：Vue3+Axios、Fetch、POST（x-www-form-urlencoded）、GET、Centos7、Mac。``

## 如何初始化项目
以当前目录 /Users/fyonecon/go/src/ 为例
```sybase
获取源代码：
cd go/src
git clone https://github.com/fyonecon/ginlaravel.git

初始化项目：
go mod init

构建依赖：
go mod tidy
go mod vendor

打包成二进制文件：
go build -mod=mod

启动MySQL数据库：请自行启动。
    
启动Redis：请自行启动。

在/config/mysql.go配置数据库信息。
在/config/redis.go配置缓存数据库信息。

启动http服务：
go run server.go

或 启动二进制文件http服务：
./ginlaravel


```

## 运行fresh热更服务（Mac环境）
以项目目录 /Users/fyonecon/go/src/ginlaravel 为例
```sybase
安装fresh：
go get github.com/pilu/fresh

去.bash_profile文件目录：
cd ~

重新编译配置文件：
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

## Swagger接口文档
官方教程：https://github.com/swaggo/gin-swagger

安装教程（GinLaravel v1.4已经集成swagger，在/routes/must.go路由文件里面，不需要再次安装和引入。）如下：
```apacheconf
1。进入项目跟目录
2。安装swag命令：
go get -u github.com/swaggo/swag/cmd/swag

3。查看是否命令是否安装成功：
swag -v

3.1。安装docs
swag init

3.2。切换到vendor目录，将swaggo安装在vendor目录下而不是全局安装：
cd vendor

4。安装 gin-swagger：
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go mod vendor

5。在route里面import引入swag：
import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs" // docs is generated by Swag CLI, you have to import it.
)

5.1。书写路由：
url := ginSwagger.URL("http://127.0.0.1:8090/swagger/doc.json") // The url pointing to API definition
route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

5.2。启动http服务：go run main.go

5.3。访问：http://127.0.0.1:8090/swagger/index.html
```

## 将项目打包成二进制文件（Mac环境）
```sybase
ginlaravel项目根目录：
go build -mod=mod

此时，在项目目录生成或更新了ginlaravel二进制文件。
开启文件的可执行权限：
chmod 773 ginlaravel
        
在项目根目录运行：
./ginlaravel
        
即可开启二进制服务。

```

## 将项目部署在Centos7上（Go环境搭建、服务器环境搭建、数据库环境搭建）
教程：https://blog.csdn.net/weixin_41827162/article/details/116048754

## 关于作者
#### 作者Author：fyonecon
##### 所在城市City：互联网不发达的长沙Changsha
##### 微信WeChat：fy66881159

## GinLaravel（GoLaravel）版权
MIT