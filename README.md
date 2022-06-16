# ReliabilityServer

#### 介绍

SDHV可靠性设计服务端程序

#### 软件架构

<img src="https://s3.bmp.ovh/imgs/2022/06/15/0e9fb66b70042ccc.png" style="zoom:50%;" />

各级目录/文件介绍

1. config: 配置目录
    1. config.yml：配置文件，包括consul节点的ip，所有节点的ip和所有的服务信息（名称，对应镜像，指令和权重），这个配置文件主要是为可靠性设计服务的
    2. serviceconfig.go: 服务模型的一些结构体
2. consul: 针对consul的一些操作
    1. consul.go: consul客户端初始化和状态检测函数，状态检测函数是关键，连续三次健康检查未通过，将该服务加入重启服务列表，并从`servicedownedtimes map`中删除，防止下次重新启动
3. healthcheck: 针对docker操作
    1. router.go: 路由
    2. handler.go: 路由对应handler, DockerOtherHandler是因为一些操作的处理逻辑相同，所以封到一个handler里了
4. param: 请求处理和发送参数
5. system: 设置reliabilityserver程序开机自启需要的文件，使用逻辑与reliabilityclient相同
6. main.go和routers.go: 与reliabilityserver逻辑相同

#### 使用说明

​	编译过程：先执行`go mod tidy`下载依赖库，然后执行`go build`，生成rs程序，直接放到ARM板上`/home/ubuntu/dang/code/reliabilityserver`目录下，然后`sudo chmod a+x rs

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

5.  https://gitee.com/gitee-stars/)