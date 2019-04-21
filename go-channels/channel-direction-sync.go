package main

import "fmt"

func putVal( in chan <-string, msg string ){
	fmt.Println("inChannel got the message")
	in <- msg
}
func  getVal(out <-chan string) string   {
	fmt.Println("get value from out channel")
	 return <- out
}
func syncChannel(in <-chan string, out chan<- string)  {
	fmt.Println("syncing the channels ")
	 out <- <- in

	 /*
	   in is an outward direction channel , so we can only take data from it .
	 	out is inward direction channel, so we can put data into it. The below statement
	 	will fail with compilation error .

	    in <- <- out
	  */
}
func main(){

	inChannel := make(chan string)
	outChannel := make(chan string)
	// inChannel <- "test"
	/*
		The above line will cause deadlock, as channel is unbuffered.
	 */
	go putVal(inChannel, "hello world ")
	go syncChannel(inChannel, outChannel)
	fmt.Println(getVal(outChannel))

	/*
		In the above 3 lines , the main will get blocked at getVal() as it is reading from the channel,
	which will make give the execution to putVal and synChannel. putVal will update the value in the inChannel.
	and gets blocked , and syncChannel with fetch the value form the inChannel and update the outChannel where it gets blocked.
	THe mail now will execute the getValue which fetches the value from the outChannel and everything gets to exit state.

	Hence when we run the program the output will be

	get value from out channel
	syncing the channels
	inChannel got the message
	hello world

	*/


}
