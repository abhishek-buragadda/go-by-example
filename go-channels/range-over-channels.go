package main

import (
	"fmt"
)

func main(){
	queue := make (chan string, 4 )

	queue <- "one"
	queue <- "two"
	close(queue) // without this line it will cause deadlock
 	for ele:= range queue {
		fmt.Println(ele)
	}

	/*
		Range iterates over all the elements as it's received in the channel.
		So if we dont close the channel then range will cause deadlock in the above scenario
		as it would be waiting for the channel to receive a value.
	 */




}
