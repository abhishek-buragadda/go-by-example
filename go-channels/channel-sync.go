package main

import (
	"fmt"
	"time"
)

func syncworker(done chan bool ){
	fmt.Println("worker started ")
	time.Sleep(time.Second )
	fmt.Println("worked ended ")
	done <- true

}
func main(){

	done := make(chan bool , 1 )
	go syncworker(done)
	fmt.Println("main started...")

	 <- done
	 /*
	 the above line will make the main thread to wait(block) and execution goes to worker
	 Once the worker is done, then the it sends message to channel and main also exit.
	  */

}