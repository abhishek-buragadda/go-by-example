package main

import (
	"fmt"
	"time"
)

func worker(  workerId int , jobs <-chan int , results chan<- int  ){
	for j := range jobs {
		val := j*j
		time.Sleep(time.Second * 2)
		fmt.Printf("\nworkerId: %d started job %d\n", workerId, j)
		results <- val
		fmt.Printf("\nworkerId: %d ended job %d\n", workerId, j)
	}

}
func main(){

	jobs:= make(chan int , 4 )
	results:= make(chan int , 4 )

	for i:=0; i< 3 ;i++  {  // total number of workers.
		go worker(i, jobs, results)
	}

	for i := 0; i < 6; i++ {  // total number of jobs created.
		jobs <- i
	}

	for i := 0; i < 6; i++ {    //  waiting for resutls from all the jobs .
		<- results
	}
	// 3 workers will handle the 6 jobs and once completed they will exit.

}

/*
	output :  There is no guarentee that worker 0 wil pick job 0 . It can go in any order based on which
goroutine will get the CPU.


workerId: 2 started job 1

workerId: 2 ended job 1

workerId: 1 started job 0

workerId: 1 ended job 0

workerId: 3 started job 2

workerId: 0 started job 3

workerId: 0 ended job 3

workerId: 3 ended job 2


 */
