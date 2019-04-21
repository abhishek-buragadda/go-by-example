package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main(){

	 var ops uint64

	for i := 0; i < 50; i++ {

		go func() {
			for{
				//atomic.AddUint64(&ops, 1)
				//ops++
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	time.Sleep(time.Second)

	sharedVal := atomic.LoadUint64(&ops)
	fmt.Println(sharedVal)
}
