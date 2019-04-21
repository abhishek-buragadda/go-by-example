package main


import (
	"fmt"
	"time"
)

func executeAfterTimerExpires( timer *time.Timer, testFun func(chan bool), done chan bool) {
	<- timer.C
	testFun(done)
	done <- true
}


func testFun(done chan bool ){
	fmt.Println("This function executed after some time")
}
func main(){
	done := make(chan bool)
	timer := time.NewTimer(time.Second* 1)
	<- timer.C
	fmt.Println("timer expired")

	newTimer := time.NewTimer(time.Second*3)
	/*
		If go is not written here, the next lines after that will get blocked until this func gets executed .
		Passing function and its arguments as function arguments in executeAfterTimeExpires.
		This is equivalent to setTimeout()
	 */
	go executeAfterTimerExpires(newTimer, testFun, done )


	/*
	We can stop an existing timer like below.

		 stopTimer:= time.NewTimer(time.Second*2)
		<- stopTimer.C
		stop := newTimer.Stop()
		if stop {
			fmt.Println("timer has be interrupted ")
		}

	 */

	fmt.Println(" not waiting for func to complete execution")


	<- done  // waiting for the goroutine to complete.

}
