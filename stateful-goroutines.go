package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)


/*
	output :

state map[3:52]
state map[3:52 7:87]
state map[0:10 3:52 7:87]
state map[0:10 3:52 5:90 7:87]
state map[0:10 2:98 3:52 5:90 7:87]
state map[0:10 2:98 3:91 5:90 7:87]
state map[0:10 2:84 3:91 5:90 7:87]
state map[0:10 2:84 3:91 5:90 7:67]
state map[0:10 2:84 3:91 5:90 7:71]
state map[0:10 2:84 3:91 4:26 5:90 7:71]
state map[0:10 2:81 3:91 4:26 5:90 7:71]
state map[0:10 2:81 3:91 4:26 5:90 7:71 9:66]
state map[0:93 2:81 3:91 4:26 5:90 7:71 9:66]
state map[0:93 2:81 3:91 4:26 5:90 7:71 9:81]
state map[0:93 2:75 3:91 4:26 5:90 7:71 9:81]
state map[0:87 2:75 3:91 4:26 5:90 7:71 9:81]
state map[0:87 2:75 3:91 4:26 5:90 7:71 9:28]
state map[0:87 2:75 3:91 4:3 5:90 7:71 9:28]
state map[0:87 2:75 3:91 4:3 5:90 6:39 7:71 9:28]
state map[0:87 1:76 2:75 3:91 4:3 5:90 6:39 7:71 9:28]
state map[0:87 1:76 2:75 3:91 4:64 5:90 6:39 7:71 9:28]
state map[0:87 1:90 2:75 3:91 4:64 5:90 6:39 7:71 9:28]
27 22

 */
func main(){
	type  read struct  {
		key int
		resp chan int
	}

	type write struct {
		key int
		val int
		resp  chan  bool
	}

	reads := make(chan *read)
	writes := make(chan *write)

	// state should be owned by the goroutine and it is only controlled via channels .
	// state should never be shared. Sharing memory by communication. :)
	go func() {
		state := make(map[int]int )
		for{
			select {
				case read:= <- reads:
					read.resp <-  state[read.key]
				case write:= <- writes:
					state[write.key] = write.val
					write.resp <- true
					fmt.Println("state", state )
			}

		}

	}()

	// create goroutines for reading
	 var totalReads uint64

	for i := 0; i < 100; i++ {
		go func() {
			for {

				read := &read{
					rand.Intn(10),
					make(chan int),
				}
				reads <- read
				<-read.resp // wait for the channel to be consumed and update the response

				atomic.AddUint64(&totalReads, 1)
				time.Sleep(100 * time.Millisecond)

			}
		}()
	}
	var totalWrites  uint64
	// create goroutines for writing
	for i := 0; i < 10; i++ {
		go func() {

			for {
				write := &write{
					rand.Intn(10),
					rand.Intn(100),
					make(chan bool),
				}
				writes <- write
				<-write.resp // wait for the channel to be consumed and update the response.
				atomic.AddUint64(&totalWrites, 1)
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 5)

	finalReadCount := atomic.LoadUint64(&totalReads)
	finalWriteCount := atomic.LoadUint64(&totalWrites)
	fmt.Println(finalReadCount, finalWriteCount)


}

