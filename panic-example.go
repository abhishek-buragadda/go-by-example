package main

import "os"

func main(){
	panic("panicking right here ")
	// all the below code wont get executed in case of panic.
	_, err := os.Create("test")
	if err!= nil {
		panic(err)
	}
}
