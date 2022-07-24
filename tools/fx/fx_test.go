package fx

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/wuqinqiang/easycar/tools"
)

func TestFrom(t *testing.T) {
	list := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, val := range list {
		var wg sync.WaitGroup
		items := val
		wg.Add(1)
		tools.GoSafe(func() {
			defer wg.Done()
			From(func(source chan<- interface{}) {
				for _, item := range items {
					source <- item
				}
			}).Walk(func(item interface{}, pipe chan<- interface{}) {
				time.Sleep(2 * time.Second)
				fmt.Println("接收参数:", item)
			}).Done()
		})
		wg.Wait()
	}

}
