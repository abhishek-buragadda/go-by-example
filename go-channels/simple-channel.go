package main

import "fmt"


// to run : go run simple-channel.go
func main(){
	messages := make(chan string )

	go func() {
		messages <- "hello world"
	}()

	msg := <- messages
	fmt.Printf( "from channel: %s", msg )

}
