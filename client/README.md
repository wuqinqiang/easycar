# easycar client for go

## install

```shell
go get -u github.com/easycar/client-go
```

## How to use

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wuqinqiang/easycar/client"
)

func main() {
	var opts []client.Option
	opts = append(opts, client.WithConnTimeout(5*time.Second))

	// new an easycar client by server uri
	cli, err := client.New("server Url", opts...)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	// begin and get gid
	gid, err := cli.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Begin gid:", gid)

	var (
		groups []*client.Group
	)
	// register  branches to easycar service (tc)
	if err = cli.Register(ctx,gid,groups); err != nil {
		log.Fatal(err)
	}
	// Trigger the execution of this distributed transaction
	if err := cli.Start(ctx,gid); err != nil {
		fmt.Println("start err:", err)
	}
	fmt.Println("end gid:", gid)
}

```

###

for more examples see here:[examples](https://github.com/easycar/examples)