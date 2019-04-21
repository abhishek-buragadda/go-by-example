package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	quit:= make(chan bool)

	go func() {
		time.Sleep(time.Second)
		//c1<- "hello"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "world"
		quit <- true
	}()

	/*
		The return statement here is like exit from the for loop.
		If we add return in any of the case, it will get exited after execution of the case.
		If we want to wait for all, then add a channel quit and add 'return' only in that case.

	 */
	for {
		select {
			case msg1 := <-c1:
				fmt.Printf("from channel 1 %s", msg1)
			case msg2 := <-c2:
				fmt.Printf("from channel2 %s", msg2)
			case <-quit:
				fmt.Println("quit called ")
				// this will exit from the for loop.
				return
		}
	}

	/*
		if you want to wait on one of the channels, remove the for loop.
		This will get exited, when  one of the channel is read.

		select {
			case msg1 := <-c1:
				fmt.Printf("from channel 1 %s", msg1)
			case msg2 := <-c2:
				fmt.Printf("from channel2 %s", msg2)
			case <-quit:
				fmt.Println("quit called ")
		}

	 */

	/*
		output :
		 time go run select-channel.go
		from channel2 worldquit called

		real    0m2.503s
		user    0m0.309s
		sys     0m0.308s


	*/



}
