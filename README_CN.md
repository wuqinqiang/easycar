# easycar:A simple distributed transaction framework implemented by go

#### easycar 是什么？

easycar 是一个用go实现的支持两阶段提交协议的分布式事务框架。它的全称是(easy commit and rollback)

更多关于easycar可以查看这篇文章->

#### Features

##### 同时支持协议和事务模式混用

在一个分布式事务中，支持每个RM协议混用(目前支持http和原生的grpc服务)，支持每个RM部分事务模式混用(目前支持TCC,Saga)。

##### 支持分层并发

支持分层并发执行。 对参与的RM通过设置的权重做分层，同一层的RM可以并发调用，一层处理完毕再接下一层。在这个基础上，当某个RM发生调用错误时，那么后面一层也不会执行，整个分布式事务需要回滚。

#### State

global state
![global](https://cdn.syst.top/global.png)

#### RUN

##### 修改配置文件
conf.yml 文件
```ymal
grpcPort: 8089
httpPort: 8085
automaticExecution2: false  #If it is true, when the first stage of execution ends, it will automatically commit or rollback
dirver: mysql
timeout: 7 #unit of second

db:
  mysql:
    dbURL: easycar:easycar@tcp(127.0.0.1:3306)/easycar?charset=utf8&parseTime=True&loc=Local
    maxLifetime: 7200
    maxIdleConns: 10
    maxOpenConns: 20
    ## add more
```

后续会提供更多配置方式


当配置完成,执行

```shell
go run cmd/main.go -mod file # mod 后续还可以是env......
```

#### examples

see more examples to:[examples](https://github.com/easycar/examples)

#### todo list

- [ ] XA
- [ ] AT
- retry
- easycar client
- more store
- ......

