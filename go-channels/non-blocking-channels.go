package main

import "fmt"

func main(){

	messages:= make(chan string )
	othermsgs := make(chan bool)


	select {
		case msg:= <- messages :
			fmt.Println(msg)
		default:
			fmt.Println("no activity")
	}

	msg := "hi"

	select {
		case messages <- msg :
			fmt.Println("this wont get printed ")
		case othermsgs <- true:
			fmt.Println("multiple case we can write, and as it is " +
				"blocking default will be executed")
		default:
			fmt.Print("no activity here also ")
	}


}
