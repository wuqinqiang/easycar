# easycar

[简体中文](https://github.com/wuqinqiang/easycar/blob/main/README_CN.md)

## What is easycar？

easycar is a distributed transaction framework implemented in go that supports a two-phase commit protocol. Its full
name is (easy commit and rollback).

For more information about easycar see this post [easycar](https://www.syst.top/posts/go/easycar/).

Architecture

![easycar](https://cdn.syst.top/easycar2.jpg)

## Features

**Supports both protocol and transaction mode mixing**

In a set of distributed transactions, each RM can use a different transport protocol (HTTP/gRPC) and transaction mode (
TCC/Sage...), so it allows a mix of RM protocols and transaction modes.

**Support for concurrent execution of transactions**

Supports concurrent execution in layers. The participating RMs are layered by the set weights, and RMs in the same layer
can be invoked concurrently, and the next layer is processed after one layer is finished.

**Service Registration and Discovery**

Currently supports etcd.

**Client-side load balancing**

**Support**：

- IPHash
- ConsistentHash
- P2C
- Random
- R2
- LeastLoad
- Bounded

## Examples of success

![success](https://cdn.syst.top/success2.png)

## Examples of failed

![success](https://cdn.syst.top/failed2.png)

## State

**global state**
![global](https://cdn.syst.top/state3.png)

## Run

```shell
cp conf.example.yml conf.ymal
```

**Modify configuration**

```ymal
## conf
automaticExecution2: true  #when the first stage of execution ends, it will commit automatically or rollback if it is true
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

db: #easycar server db
  driver: mongodb
  mysql:
    dbURL: easycar:easycar@tcp(127.0.0.1:3306)/easycar?charset=utf8&parseTime=True&loc=Local
    maxLifetime: 7200
    maxIdleConns: 10
    maxOpenConns: 20
  mongodb:
    url: mongodb://127.0.0.1:27017/easycar
    minPool: 10
    maxPool: 20

registry: 
  etcd:
    user: ""
    pass: ""
    hosts:
      - 127.0.0.1:2379
  #add more

tracing:
  jaegerUrl: http://localhost:14268/api/traces

cron:
  maxTimes: 2   #max retry times when rm is not available
  timeInterval: 1 #unit is minute. it means that the next retry is 1m later, not in strict mode  
```

**run**
```shell
go run cmd/main.go 
```

## examples

more examples to:[examples](https://github.com/easycar/examples)

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
    - [x] consul
- [x] balancer
- [ ] test
- [ ] notify
- [x] tracing
- [ ] tool
- more store
- ......

