package main

import (
	"fmt"
	"sync"
)

type MyCounter struct {
	mx sync.Mutex//встраиваем мьютекс в структуру, так как счетчик должен быть потокобезопасным
	c  int
}

func counter(wg *sync.WaitGroup, m *MyCounter) {
	m.mx.Lock()
	m.c++
	m.mx.Unlock()
	wg.Done()
}

func main() {
	n := 200
	var c MyCounter
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go counter(&wg, &c)
	}
	wg.Wait()

	fmt.Println(c.c)
}
