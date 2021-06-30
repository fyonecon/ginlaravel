
########### Dockerfile代码 ##########

#1.1）docker images查看制作
FROM golang:latest as builder
# 修改系统为上海时区
RUN echo "Asia/Shanghai" > /etc/timezone \
 && rm /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
#作者
MAINTAINER fyonecon "youcansendmsg@qq.com"
#设置工作目录（暂时规定最终文件夹名与module名相同）
WORKDIR $GOPATH/src/ginvel
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/ginvel
#修改env参数
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN export GIN_MODE=release
#初始化框架所需扩展
# RUN go mod vendor
#编译成二进制文件
RUN go build main.go
#RUN go build -o /tmp/main main.go
#暴露端口
EXPOSE 8090
#最终运行docker的命令（运行项目生成的二进制文件）
ENTRYPOINT  ["./main"]

#FROM debian:slim
#WORKDIR /go/bin
#COPY --from=builder /tmp/main ./main
#ENTRYPOINT ["/go/bin/main"]

######### Dockerfile教程 #########

#1.2）制作docker镜像，单独在项目跟目录运行（这个docker镜像名可以自定义，一般与文件夹名或module名一样即可）：
#将Dockfile文件放置在项目跟目录，然后运行
#docker build -t ginvel.com .
#查看镜像
#docker images
#查看正在运行的docker
#docker ps
#查看docker占用内存大小
#docker ps -as

#2）docker运行镜像，运行生成的二进制文件（二进制文件为项目go.mod里面的module名，不可更改。及其重启后，docker的端口映射需要重新运行。）
#第一个端口为宿主机端口，第二个端口为docker端口，注意httpserver的host是0.0.0.0。
#开启docker（前台方式）
#docker run -p 8090:8090 ginvel.com
#开启docker（后台方式，推荐）
#docker run -p 8090:8090 -d ginvel.com

#2.1）查看已运行的docker
#docker ps

#3）查看docker日志
#docker logs -f ginvel.com

#4）其他（可以使用docker build的名或镜像id，用docker images查看）
#停止服务
#docker stop ginvel.com
#删除container实例
#docker rm ginvel.com
#删除container镜像
#docker rmi ginvel.com
#强行删除container镜像
#docker rmi -f ginvel.com

#5）进入docker的命令
#查看正在运行的docker的id（docker images不行）
#sudo docker ps
#进入docker
#sudo docker exec -it 正在运行的docker的id /bin/bash


