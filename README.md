# easycar

what is easycar?

a simple Distributed transactions implemented by go. the full name of easycar is (easy commit and rollback).

### role

- coordinator
- TM
- RM

### plan

### struct

```go
type Branch struct {
gId               string
url               string
reqData           string
branchId          string
PId  string     // 父级事务
respData string //分支事务执行结果，子事务依赖父级事务的结果
}
transactionAction consts.BranchAction
state             consts.BranchState
protocol string //http or grpc
endTime  int64
}


type Global struct {
gId      string
state    consts.GlobalState
endTime int64
}
```