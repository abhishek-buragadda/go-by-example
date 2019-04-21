package main

import (
	"fmt"
	"time"
)

func main (){
	c1 := make(chan string )
	c2 := make (chan string )

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "channel 1 "
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "channel 2"
	}()

	select {
		case m1:= <- c1 :
			fmt.Printf("channel 1 %s ", m1)
		case <- time.After(time.Second *1):
			fmt.Println("timeout for c1")
	}

	select {
		case m2:= <- c2 :
			fmt.Printf("channel2  %s", m2)
		case <- time.After(time.Second * 4):
			fmt.Println("timeout for c2")
	}
}
