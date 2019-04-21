package main

import "fmt"

func main(){

	/*
		Observations :
		 - Order is always preserved in go channel
	     - If more messages are inserted into the channel than its capacity, it is sleeping and r
			program is deadlocked.Find out why and how?
		-

	 */
	stringChannel := make(chan string , 2)

	stringChannel <- "test"
	stringChannel <- "test2"

	allChannel := make(chan interface{}, 2)
	allChannel <- 23
	allChannel <- "asdfdsaf"


	fmt.Println(<- allChannel)
	fmt.Println( <- stringChannel)
	allChannel <- "test1245"
	fmt.Println( <-allChannel )




}
