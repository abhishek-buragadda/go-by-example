package main

import (
	"fmt"
	"time"
)

func main(){

	requests := make (chan int , 5 )
	for i := 0; i < 5; i++ {
		requests <- i
	}

	close(requests)

	/*
		limiting the request to every 200ms even though we get requests at a higher rate.
	 */
	limiter := time.Tick(time.Millisecond * 200)
	for r := range requests {
		<- limiter
		fmt.Println("request", r , time.Now())
	}

	/*
		allowing some number of requests at a time like 3 requests are allowed at one and once they are
		done we will allow the next 3.
	 */
	burstLimiter := make(chan time.Time,3 )
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}
	go func() {
		for   t := range time.Tick(time.Second){
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5 )
	for i := 0; i < 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for r := range burstRequests {
		<- burstLimiter
		fmt.Println("burst request:", r , time.Now())
	}

}