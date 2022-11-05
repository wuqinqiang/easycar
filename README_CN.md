# easycar

## easycar 是什么？

easycar 是一个用go实现的支持两阶段提交协议的分布式事务框架。它的全称是(easy commit and rollback).

更多关于easycar看这篇文章:[easycar](https://www.syst.top/posts/go/easycar/)

## 架构图

![easycar](https://cdn.syst.top/easycar.png)

## Features

**同时支持协议和事务模式混用**

在一组分布式事务中，每个RM可以使用不同的传输协议(HTTP/gRPC),也可以使用不同的事务模式(TCC/Sage...)，因此允许RM协议和事务模式的混合使用。

**支持并发执行事务**

支持分层并发执行每个RM。 对参与的RM设置分层，同一层的RM可以并发调用，一层处理完毕再接下一层。

**服务注册和发现**

暂时只支持etcd。

**负责均衡**

提供：

- IPHash
- ConsistentHash
- P2C
- Random
- R2
- LeastLoad
- Bounded


## 成功的例子

![success](https://cdn.syst.top/success.png)

## 失败的例子
![failed](https://cdn.syst.top/failed.png)

## 状态

![global](https://cdn.syst.top/b-state.png)

## 运行

```shell
cp conf.example.yml conf.ymal
```

修改 conf.yml 文件

```ymal
## conf
#httpListen: 127.0.0.1:8085
automaticExecution2: false  # 如果是true，easycar 会自动根据第一阶段结果，执行第二阶段commit或者rollback
timeout: 7 #unit of second
server:
  grpc:
    listenOn: 127.0.0.1:8088
    keyFile:   #server key
    certFile:  #server cert
    gateway:
      isOpen: true
      certFile:  #client cert
      serverName:
  http:
    listenOn: 127.0.0.1:8085

db: # easycar server db
  driver: mysql
  mysql:
    dbURL: easycar:easycar@tcp(127.0.0.1:3306)/easycar?charset=utf8&parseTime=True&loc=Local
    maxLifetime: 7200
    maxIdleConns: 10
    maxOpenConns: 20
  mongodb:
    url: mongodb://127.0.0.1:27017/easycar
    minPool: 10
    maxPool: 20

registry: #//配置了注册中心，那么服务启动的时候把服务注册到注册中心
  etcd:
    user: ""
    pass: ""
    hosts:
      - 127.0.0.1:2379
  ## add more

tracing:
  jaegerUrl: http://localhost:14268/api/traces
```

执行

```shell
go run cmd/main.go
```

## client

如果你使用Golang，可以使用 [client](https://github.com/easycar/client-go) ,其他语言客户端后续提供。

## examples

更多例子:[examples](https://github.com/easycar/examples)

## todo list

- [x] Saga
- [x] TCC
- [ ] XA
- [ ] client
    - [x] client-go
    - [ ] client-rust
    - [ ] client-php
    - [ ] client-python
    - [ ] client-java
- [x] retry
- [ ] registry and discovery
    - [x] etcd
    - [ ] consul
    - [ ]  zookeeper
- [x] balancer
- [ ] test
- [ ] notify
- [x] tracing
- [ ] tool
- more store
- ......

