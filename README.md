## 欢迎 Star🌟 本项目
#### Github（推荐）：https://github.com/fyonecon/ginlaravel

## 版本
#### v1.9

## 简介
GinLaravel基于Golang框架Gin（Gin学习文档：https://learnku.com/docs/gin-gonic/2019 ）开发，项目结构和项目理念参考Laravel。学习本项目时建议从阅读本项目源码开始，并具备Golang、Gin、go mod、centos+nginx、redis、mysql、加密解密、http、Swagger等必要知识。GinLaravel的构建和运行周期与Beego、Vue、React、Laravel、ThinkPHP、Django等都会有类似的引导思路、参数设置、插件扩展、服务部署、代码统一性、生态护城河等。GinLaravel尽量减少汉字造词、功能造词、贤者造词，不会用描述华丽、中英文混杂的造词方式去包装一个功能平常的老东西（比如 http+header、cookie/localstorage/sqlite+token、加密与解密、数据缓存+队列、负载均衡+docker、内存常驻与微服务、接口安全与前后端分离等）。

整个开发以人为本，功能模块化扩展，系统安全可控，渐进式提升访问量。

GinLaravel支持MVC开发模式。本项目展示了MVC处理数据，同时展示"Verify—Controller—Kit"模式处理数据。

## 现已支持
> Go-MySQL、GORM（v2）、Go-Redis、热更（fresh）、Swagger、MVC、模版输出（v1.9开始已剔除）、Http访问频率拦截、熔断、HttpCors、对称加密（可中文）、http拦截器、多层路由、分词与全文检索、运行时监控、图形验证码、生成和读取Excel、全局定时器等。

> 测试过的客户端环境等：Vue3+Axios、Fetch、POST（x-www-form-urlencoded）、GET、Centos7、Mac。

## 设计理念
> 宽进严出，面向Api，适合复杂项目，任何参数或服务都会有默认值。整个项目运行需go+mysql+redis！

## 路由周期
> 请求路由名 ➡️ header过滤 ➡️ 拦截请求频率 ➡️ 熔断 ➡️ 校验请求方法和Token参数 ➡️ 运行目标函数 ➡️ Controller程序并返回Json ➡️ 治理或运行时监控服务 ➡️ 程序结束️

## 项目目录
+ /app/Common/ ※ 对go原生封装的公共函数、自定义必要数据参数配置。有很多用Go解释Go的公用函数。

+ /app/Http/ ※ 控制器、模型、拦截器

+ /app/Http/Controllers/ ※ 控制器

+ /app/Http/Controllers/Example ※ 示例

+ /app/Http/Controllers/Gen3 ※ 版本3的控制器文件夹

+ /app/Http/Models/ ※ 数据模型

+ /app/Http/Middlewares/ ※ 中间件，含有cors、http限速、500报错拦截、默认路由、app运行时等

+ /app/Kit/ ※ 自定义系统服务，包含第三方服务和系统功能服务

+ /app/Ruler/Runtime/ ※ app服务等运行时事件处理

+ /app/Ruler/Task/ ※ app全局定时任务，默认20s精度

+ /bootstrap/ ※ 系统服务启动、数据库全局引用配置，一般不需要更改此处。

+ /bootstrap/app.go ※ 项目运行入口

+ /bootstrap/init.go ※ 项目全局参数初始化

+ /bootstrap/driver/ ※ 服务驱动

+ /config/ ※ MySQL数据库、Redis缓存、http端口号等配置

+ /config/docs/ ※ 接口文档插件Swagger的目录

+ /docs/ ※ 自动接口文档Swagger的的目录

+ /extend/ ※ 自定义的扩展库，kit是扩展接口应用，extend是扩展接口封装

+ /routes/ ※ 路由，同时支持模版型路由（web.go）、Api型路由（api.go）。推荐4层路由命名。

+ /tmp/ ※ fresh热更的缓存日志目录。目录需权限777。v1.9已经迁移到storage文件夹里面的fresh_tmp文件夹。

+ /storage/ ※ 系统日志、文件上传、静态缓存。目录需权限777。

+ /views/ ※ 模版渲染的原始文件夹（v1.9已剔除）

+ /views/pages/ ※ 模版的html文件（v1.9已剔除）

+ /views/static/ ※ 模板静态资源：js、css、img（v1.9已迁移）

+ ginvel.com ※ 项目生产的二进制文件，在生产环境用。目录需权限773。

+ go.mod ※ 项目所依赖的module路径、第三方库等的引入

+ main.go ※ 跑起本项目的入口main文件。

## 命名原则
+ 自定义函数：大驼峰
+ 自定义变量：小驼峰
+ 自定义结构体和结构体成员：大驼峰
+ MySQL、Redis：小写+下划线
+ 接口名：小写+下划线
+ 自定义文件夹名：大驼峰
+ 系统集文件夹名：小写

## 基础环境
> 1⃣️Go运行环境
> 

> 2⃣️MySQL5.7+
> 
> 请将/项目资料/ginlaravel.sql 文件导入到你的数据库

> 3⃣️Redis
>
> 请提前开启你的Redis服务

> cmd中运行「 go run main.go 」即可启动项目。
> 
> 或使用热更方式启动http服务，在cmd中目录运行"fresh"。v1.4开始为适配swaggo遂将server.go更名为main.go。

> 访问
> 
> http://127.0.0.1:8090

> 项目上线
> 
> serverConfig["ENV"]的值改成release，然后使用以上同样方法运行。

## 如何初始化项目
以当前目录 /Users/fyonecon/go/src/ 为例
```cmd
获取源代码：
cd go/src
git clone https://github.com/fyonecon/ginlaravel.git

初始化项目：
go mod init

清除不需要的vendor中的第三方扩展：
go mod tidy
    
重新载入vendor中的第三方扩展：
go mod vendor

将项目打包成二进制文件：
go build -mod=mod
（运行二进制文件需要ginlaravel的文件的权限为：chmod 773 ginvel.com）

在/config/mysql.go配置数据库信息。
在/config/redis.go配置缓存数据库信息。

启动http服务：
go run main.go

或 启动二进制文件http服务：
./ginlaravel


```

[comment]: <> (## Go包管理工具Vendor)

[comment]: <> (```apacheconf)

[comment]: <> (下载govendor工具到本地：)

[comment]: <> (go get -u github.com/kardianos/govendor)

[comment]: <> (生成vendor文件夹（存放你项目需要的依赖包）和vendor.json：)

[comment]: <> (govendor init)

[comment]: <> (将GOPATH文件夹中的包添加到vendor目录下：)

[comment]: <> (govendor add +external)

[comment]: <> (或者govendor add +e)

[comment]: <> (或者govendor add +e)

[comment]: <> (govendor update +vendor # 更新vendor的包命令)

[comment]: <> (govendor list)

[comment]: <> (# govendor fetch github.com/xxx)

[comment]: <> (```)

## 运行fresh热更服务（Mac环境）
以项目目录 /Users/fyonecon/go/src/ginlaravel 为例
```cmd
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

> 热更服务文档：https://github.com/gravityblast/fresh 。

## Swagger接口文档
>官方教程：https://github.com/swaggo/gin-swagger

安装教程（GinLaravel v1.4已经集成swagger(v1.7开始默认不集成，因为没什么卵用)，在/routes/must.go路由文件里面，不需要再次安装和引入。）如下：
```cmd
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
	_ "ginlaravel/docs" // docs is generated by Swag CLI, you have to import it.
)

5.1。书写路由：
url := ginSwagger.URL("http://127.0.0.1:8090/swagger/doc.json") // The url pointing to API definition
route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

5.2。启动http服务：go run main.go

5.3。访问：http://127.0.0.1:8090/swagger/index.html
```

## 将项目打包成二进制文件（Mac、Centos环境）
```cmd
1。在ginlaravel项目根目录：
go build -mod=mod

此时，在项目目录生成或更新了ginlaravel二进制文件。
    
2。开启文件的可执行权限：
chmod 773 ginlaravel
        
3。在项目根目录运行：
./ginvel.com
        
即可开启二进制服务。

```

## 将Go项目部署在Centos7上（Go环境搭建、服务器环境搭建、数据库环境搭建）
> 教程：https://blog.csdn.net/weixin_41827162/article/details/116048754

## 如何将GOPATH里面的"go get xxx"扩展引用到vendor里面
> 先运行"go get 扩展github地址"，在项目中import引入"xxx"扩展
> 
> 再在项目跟目录运行"go mod vendor"
> 
> 这样就可以将扩展自动引用到vendor目录下而不用govendor。

## 作者Author
> https://github.com/fyonecon

## GinLaravel（Ginvel）版权
> MIT（可更可改可匿名可传播，商用不受限）
