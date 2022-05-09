package main

import (
	"fmt"
	//"time"
)

func first(chan_1 chan int){
	for {
        _, ok := <-chan_1
        if ok == false {
            break // exit break loop
        } else {
            
        }
    }
}

func second(chan_2 chan int){
	for {
        val, ok := <-chan_2
        if ok == false {
            break // exit break loop
        } else {
            fmt.Println(val)
        }
    }
	
}

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	chan_1:= make(chan int)
	chan_2:= make(chan int)
	go first(chan_1)
	go second(chan_2)
	for _, v := range pow {
		chan_1<-v
		chan_2<-v*2
		//time.Sleep(time.Nanosecond)
	}
	close(chan_1)
	close(chan_2)
}