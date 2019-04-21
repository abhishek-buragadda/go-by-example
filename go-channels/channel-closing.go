package main

import "fmt"

func main(){


	jobs := make(chan int)
	done :=  make(chan bool )
	test := make(chan string)
	go func() {
		for{
			j, more := <- jobs
			if more {
				fmt.Println("after received job ", j )
			} else{
				fmt.Println("received all jobs")
				test <- "test"
				done <- true
				close(test)
				return
			}
		}

	}()

	for  i := 0; i <= 3; i++ {
		fmt.Println("before job is  pushed ", i)
		jobs <- i
		fmt.Println("after job is  pushed ", i)
	}

	close(jobs)
	fmt.Println("printing from closed channel", <-test)
	fmt.Println("closed all jobs ")

	<- done


/*

	THis is an unbufferred channel so jobs should be removed before new one has to be pushed  or else that execution would be blocked.
	THe same is  seen in the output . Before the 1 is pushed 0 has to be consumed. The main thread will be blocked and once 0 is consumed
	the main is allowed to push the next job 2 .

	If we have a buffered channel, the channel will get filled until the buffer and then the same behavior as above is seen after that .

	Closed channels  :
		Once  a channel is closed we cannot send anything in the channel. It will be a error.
		You can still receive from a closed channel.
	output :

	before job is  pushed  0
   after job is  pushed  0
   before job is  pushed  1
   after received job  0
   after received job  1
   after job is  pushed  1
   before job is  pushed  2
   after job is  pushed  2
   before job is  pushed  3
   after received job  2
   after received job  3
   after job is  pushed  3
   closed all jobs
   received all jobs

 */

}
