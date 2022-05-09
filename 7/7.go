package main

import (
	"fmt"
	"sync"
)

type MyMap struct {//структура мап с мьютексом, который позволяет провести конкурентную запись в мап, так как мап для записи не потокобезопасен
	mx sync.Mutex
	m map[int]int
}

func NewMap() *MyMap {
    return &MyMap{
        m: make(map[int]int),
    }
}

func (c *MyMap) Load(key int) (int, bool) {
    c.mx.Lock()
    defer c.mx.Unlock()
    val, ok := c.m[key]
    return val, ok
}

func (c *MyMap) Store(key int, value int) {//запись в мап
    c.mx.Lock()
    defer c.mx.Unlock()
    c.m[key] = value
}

func job(jobs chan int, mymap *MyMap){
	a:=15
	for {
		mymap.Store(a, <-jobs)
		if v, ok := mymap.Load(a); ok == false {
			fmt.Println("Something went wrong")
		} else {
			fmt.Println(a," ", v)
		}
		a++
	}
}

func main() {
	jobs := make(chan int)
	mymap := NewMap()
	go job(jobs, mymap)
	for i := 0; i < 10; i++ {
		jobs <- i
	}
	fmt.Println(mymap)

}