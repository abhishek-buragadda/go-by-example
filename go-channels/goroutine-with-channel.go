package main


import "fmt"

func greet(c chan string, wait chan string  ){
	fmt.Println(<- c)
	fmt.Println(<- c)
	wait <- "done"
}

func main(){

	waitChannel := make(chan string)    // unbuffered channel
	/* Unbuffered channel will wait/block  until the value is consumed.
		So the below code will result in deadlock as we cannot put another value in the
		unbuffered channel without consuming it .

		waitChannel <- "test "
		waitChannel <- "new test "

	 */
	channel  := make(chan string, 2 )  // buffered channel
	channel <- "Hello"
	go greet(channel, waitChannel)
	channel <- "world"

	<- waitChannel
	fmt.Println("channel empty now ..")
	//close(channel)



}