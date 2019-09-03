## centos 配置go-mrico框架（grpc，protoc，protoc-gen-go，protoc-mrico-go，consul）

> 最近转行当go的小弟，然后开始尝试 golang微服务框架 go-mrico ，结果才发现这东西真的鸡儿难装，搞了几天终于成功之后，好好总结下。
>  测试项目地址 https://github.com/oldguys/MyMircoDemo
>


**go-mrico 官网： [https://micro.mu/docs/go-micro.html](https://micro.mu/docs/go-micro.html)**


### 安装步骤:
1. 安装配置 golang 运行环境
2. 配置 运行时需要的 protoc
3. 安装 go-mrico 通讯依赖 grpc
4. 安装 go-mrico 以及相关依赖
5. 安装 **\*.mrico.go 和 \*.pb.go**  代码生成工具
6. 安装微服务注册中心 consul


安装centos

使用 ip addr 查看ip

### 1. 配置 golang 运行环境

官网下载 [https://golang.google.cn/dl/](https://golang.google.cn/dl/)

下载安装包 go1.12.7.linux-amd64.tar.gz

cd 到安装包位置

tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz

**配置 golang 环境变量**

***PS：使用 export PATH=$PATH:/usr/local/go/bin 重启系统会失效，所以需要使用以下环境变量配置的方式***
```
cd /etc/profile.d
```
使用 vim golang.sh 创建 golang.sh



```
export GOROOT=/usr/local/go
export GOPATH=/root/go  
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin:$GOBIN
```
__重启系统 (不要忘了!)__

![验证效果](https://upload-images.jianshu.io/upload_images/14387783-020114a1251717e5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 2.  配置 运行时需要的 protoc

下载地址  https://github.com/protocolbuffers/protobuf/releases  安装包：protobuf-all-3.9.1.zip

安装 unzip 工具
```
yum install -y unzip
```
将压缩包移动到 /usr/local/ ，并解压
将名称更改为protobuf
```
mv /usr/temp/protobuf-all-3.9.1.zip  /usr/local/
unzip protobuf-all-3.9.1.zip
mv protobuf-3.9.1/ protobuf
```
**安装 make**
```
yum -y install gcc automake autoconf libtool make
yum install gcc gcc-c++
```

**配置make**
```
cd protobuf
./configure --prefix=/usr/local/protobuf

make
# 测试，这一步很耗时间
make check
make install
```
![002.png](https://upload-images.jianshu.io/upload_images/14387783-c67d2dcb0bd4c920.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

同上配置环境变量
```
cd /etc/profile.d
```
创建proto.sh
```
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/protobuf/lib

export LIBRARY_PATH=$LIBRARY_PATH:/usr/local/protobuf/lib

export PATH=$PATH:/usr/local/protobuf/bin
```
__重启系统 (不要忘了!)__，验证是否安装成功

![003.png](https://upload-images.jianshu.io/upload_images/14387783-181591397c080c29.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 3.  安装go-mrico 通讯依赖 grpc

官网 QuickStart [https://grpc.io/docs/quickstart/go/](https://grpc.io/docs/quickstart/go/)

git地址 [https://github.com/grpc/grpc-go](https://github.com/grpc/grpc-go)

默认使用，可以下载相关的依赖，但是实际特别漫长，而且网络经常出问题
```
go get -u google.golang.org/grpc
```
所以可以使用 git下载替代（以下直使用git下载方式）
```
## 初始化 git 安装工具，已有可缺省
yum install -y git
## download go-mrico相关依赖
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
## 代码生成器
git clone https://github.com/golang/protobuf.git /root/go/src/github.com/golang/protobuf
```
验证效果
```
## 测试用例位置
cd $GOPATH/src/google.golang.org/grpc/examples/helloworld
### 服务端
go run greeter_server/main.go
### 客户机端
go run greeter_client/main.go
```
如果抛出异常，如图
![004.png](https://upload-images.jianshu.io/upload_images/14387783-b6ebec5d075d4ef6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
原因是缺少依赖，补全依赖即可
```
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/googleapis/go-genproto.git $GOPATH/src/google.golang.org/genproto
```
运行结果：
![005.png](https://upload-images.jianshu.io/upload_images/14387783-f3fc58b8aa0e946f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![006.png](https://upload-images.jianshu.io/upload_images/14387783-5c9f24712ed5cf55.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 4.  安装go-mrico 以及相关依赖

默认使用，可以下载相关的依赖，但是实际特别漫长，而且网络经常出问题
```
  go get -u -v github.com/micro/micro
```
同样的，还是使用git下载快稳一些
```
yum install -y git
## download go-mrico相关依赖
git clone https://github.com/micro/micro.git  $GOPATH/src/github.com/micro/micro
```
尝试安装
```
go install github.com/micro/micro
```
同样出现上面的结果
![007.png](https://upload-images.jianshu.io/upload_images/14387783-2e657a98a4f7fbff.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

同理，从git上一个个下载下来（这边就比maven不人性化多了，当然也可能因为有阿里镜像，这边没有）规模有点浩大
```
git clone https://github.com/google/btree.git $GOPATH/src/github.com/google/btree

git clone https://github.com/pkg/errors.git $GOPATH/src/github.com/pkg/errors

git clone https://github.com/armon/go-metrics.git  $GOPATH/src/github.com/armon/go-metrics

git clone https://github.com/gorilla/handlers.git $GOPATH/src/github.com/gorilla/handlers

git clone https://github.com/gorilla/mux.git $GOPATH/src/github.com/gorilla/mux

git clone https://github.com/hashicorp/go-immutable-radix.git  $GOPATH/src/github.com/hashicorp/go-immutable-radix

git clone https://github.com/hashicorp/go-msgpack.git  $GOPATH/src/github.com/hashicorp/go-msgpack

git clone https://github.com/hashicorp/go-multierror.git  $GOPATH/src/github.com/hashicorp/go-multierror

git clone https://github.com/hashicorp/golang-lru.git $GOPATH/src/github.com/hashicorp/golang-lru

git clone https://github.com/hashicorp/errwrap.git $GOPATH/src/github.com/hashicorp/errwrap

git clone https://github.com/hashicorp/go-sockaddr.git $GOPATH/src/github.com/hashicorp/go-sockaddr

git clone https://github.com/joncalhoun/qson.git $GOPATH/src/github.com/joncalhoun/qson

git clone https://github.com/lucas-clemente/quic-go.git $GOPATH/src/github.com/lucas-clemente/quic-go

git clone https://github.com/cheekybits/genny.git $GOPATH/src/github.com/cheekybits/genny

git clone https://github.com/marten-seemann/qtls.git $GOPATH/src/github.com/marten-seemann/qtls

git clone https://github.com/nlopes/slack.git $GOPATH/src/github.com/nlopes/slack

git clone https://github.com/sean-/seed.git $GOPATH/src/github.com/sean-/seed

git clone https://github.com/serenize/snaker.git  $GOPATH/src/github.com/serenize/snaker

git clone https://github.com/xlab/treeprint.git  $GOPATH/src/github.com/xlab/treeprint

git clone http://www.github.com/go-telegram-bot-api/telegram-bot-api $GOPATH/src/gopkg.in/telegram-bot-api.v4

git clone https://github.com/technoweenie/multipartstreamer.git $GOPATH/src/github.com/technoweenie/multipartstreamer

git clone https://github.com/micro/go-micro.git $GOPATH/src/github.com/micro/go-micro

git clone https://github.com/go-log/log.git $GOPATH/src/github.com/go-log/log

git clone https://github.com/google/uuid.git  $GOPATH/src/github.com/google/uuid

git clone https://github.com/hashicorp/consul.git $GOPATH/src/github.com/hashicorp/consul

git clone https://github.com/hashicorp/memberlist.git  $GOPATH/src/github.com/hashicorp/memberlist

git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto

git clone https://github.com/micro/mdns.git $GOPATH/src/github.com/micro/mdns

git clone https://github.com/micro/cli.git $GOPATH/src/github.com/micro/cli

git clone https://github.com/json-iterator/go.git $GOPATH/src/github.com/json-iterator/go

git clone https://github.com/miekg/dns.git $GOPATH/src/github.com/miekg/dns

git clone https://github.com/mitchellh/hashstructure.git  $GOPATH/src/github.com/mitchellh/hashstructure

git clone https://github.com/modern-go/concurrent.git $GOPATH/src/github.com/modern-go/concurrent

git clone https://github.com/modern-go/reflect2.git $GOPATH/src/github.com/modern-go/reflect2

git clone https://github.com/nats-io/nats.go.git $GOPATH/src/github.com/nats-io/nats.go

git clone https://github.com/nats-io/jwt.git  $GOPATH/src/github.com/nats-io/jwt

git clone https://github.com/nats-io/nkeys.git $GOPATH/src/github.com/nats-io/nkeys

git clone https://github.com/nats-io/nuid.git $GOPATH/src/github.com/nats-io/nuid

git clone https://github.com/bwmarrin/discordgo.git $GOPATH/src/github.com/bwmarrin/discordgo

git clone https://github.com/chzyer/readline.git $GOPATH/src/github.com/chzyer/readline

git clone https://github.com/forestgiant/sliceutil.git $GOPATH/src/github.com/forestgiant/sliceutil

git clone https://github.com/gorilla/websocket.git $GOPATH/src/github.com/gorilla/websocket
```
安装成功，验证版本
```
micro --version
## 查看是否在bin目录下 生成可执行文件
ls $GOPATH/bin/
```
![008.png](https://upload-images.jianshu.io/upload_images/14387783-7a0554c388093fb0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 5.  安装go-mrico 代码生成工具

github地址： [https://github.com/micro/protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

```
git clone https://github.com/micro/protoc-gen-micro.git $GOPATH/src/github.com/micro/protoc-gen-micro

go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go
## 验证安装成功
ls $GOPATH/bin/

```
![009.png](https://upload-images.jianshu.io/upload_images/14387783-ac6de6a63d8e2d77.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

进行测试
```

cd $GOPATH/src/github.com/micro/protoc-gen-micro/examples/greeter
# 编译生成
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```
效果图
![010.png](https://upload-images.jianshu.io/upload_images/14387783-41dde11419e7d9b9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 6. 安装微服务注册中心 consul

官网： [https://www.consul.io/docs/install/index.html#compiling-from-source](https://www.consul.io/docs/install/index.html#compiling-from-source)

会出现缺少 依赖，补充依赖，提前补充下
![012.png](https://upload-images.jianshu.io/upload_images/14387783-fe73ea7dd1c3a036.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```
git clone https://github.com/golang/tools.git  $GOPATH/src/golang.org/x/tools
```
正式安装
```
cd $GOPATH/src/github.com/hashicorp
### 如果已经存在 可以缺省
git clone https://github.com/hashicorp/consul.git

cd consul/
## 开始安装，特别漫长
make tools

make dev

consul -v
```
效果图：

![013.png](https://upload-images.jianshu.io/upload_images/14387783-ab6b91ddce1e7bb6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

以上基本完成安装 ( 撒花 ~）
### 最后开始最终测试 :

官网用例： [https://micro.mu/docs/go-micro.html](https://micro.mu/docs/go-micro.html)

```
git clone https://github.com/micro/examples.git $GOPATH/src/github.com/micro/examples
## 启动
consul agent -dev

#进入 测试代码目录 
cd $GOPATH/src/github.com/micro/examples/service

## 启动服务端
go run main.go --registry=consul

## 启动客户端
go run main.go --run_client  --registry=consul
```
效果图
![014.png](https://upload-images.jianshu.io/upload_images/14387783-9cb746208ee5fbb0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![015.png](https://upload-images.jianshu.io/upload_images/14387783-0479fcde9626b0ff.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![016.png](https://upload-images.jianshu.io/upload_images/14387783-c787e2c5df5ad887.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 同样的 window 也可以这样配置，可以直接把这边下载好的 依赖直接复制到window环境中，可以省超级多时间，具体操作就省略了。
直接贴效果图

![017.png](https://upload-images.jianshu.io/upload_images/14387783-869117b87ed6cae0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![018.png](https://upload-images.jianshu.io/upload_images/14387783-dc782ae4692a98d6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![019.png](https://upload-images.jianshu.io/upload_images/14387783-4d0f2d09022e2955.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**注意运行时配置 运行时参数**
![020.png](https://upload-images.jianshu.io/upload_images/14387783-9a8444697dfbb5a5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
