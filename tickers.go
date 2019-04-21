package main

import (
	"fmt"
	"time"
)

func main(){

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C  {   // this will continously give values every 500ms
			// execute any task after the ticker time.
			fmt.Println(t)
		}
	}()
	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("Done with ticker")
}