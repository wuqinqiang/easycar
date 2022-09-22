# easycar:A simple distributed transaction framework implemented by go

[简体中文](https://github.com/wuqinqiang/easycar/blob/main/README_CN.md)
## What is easycar？

easycar is a distributed transaction framework implemented in go that supports a two-phase commit protocol. Its full name is (easy commit and rollback).

### Features

#### Supports both protocol and transaction mode mixing

Support for mixed use of each RM protocol in a distributed transaction (currently supports http and native grpc services). Support per RM transaction mode mix.

#### Support for concurrent execution of transactions

Supports concurrent execution in layers. The participating RMs are layered by the set weights, and RMs in the same layer can be invoked concurrently, and the next layer is processed after one layer is finished. On this basis, when a RM has a call error, then the next layer will not be executed and the whole distributed transaction needs to be rolled back.





More about easycar can check this article 



## State

global state
![global](https://cdn.syst.top/global.png)

## RUN

### Modify configuration
conf.yml file
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

More configuration methods will be provided later


When the configuration is complete, execute

```shell
go run cmd/main.go -mod file # The follow-up can also be env、etcd......
```
If you use golang,use [client](https://github.com/easycar/client-go).
of course, you can use directly http api.



## examples

see more examples to:[examples](https://github.com/easycar/examples)

## todo list
- [ ] XA
- [ ] AT
- retry
- easycar client
- more store
- test
- ......

