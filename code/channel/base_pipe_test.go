package channel

import (
	"fmt"
	"testing"
	"time"
)

//求偶数
func Event(args ...interface{}) InChan {
	c := make(chan interface{})
	go func() {
		defer close(c)
		for _, num := range args {
			if num.(int)%2 == 0 {
				c <- num
			}
		}
	}()
	return c
}

//乘以10
func M10(in InChan) OutChan { //这个函数是支持管道的
	out := make(chan interface{})
	go func() {
		defer close(out)
		for num := range in {
			time.Sleep(time.Second * 1)
			out <- num.(int) * 10
		}
	}()
	return out
}

func TestNewPipe(t *testing.T) {
	nums := []int{2, 3, 6, 12, 22, 16, 4, 9, 23, 64, 62}
	interfaceSlice := make([]interface{}, len(nums))
	for i, d := range nums {
		interfaceSlice[i] = d
	}
	start := time.Now().Unix()
	pipe := NewPipe()
	pipe.SetCmd(Event)
	pipe.SetPipeCmd(M10, 8)
	ret := pipe.Exec(interfaceSlice...)
	end := time.Now().Unix()
	t.Logf("测试--用时:%d秒\r\n", end-start)
	for r := range ret {
		t.Logf("%d ", r)
	}
}
