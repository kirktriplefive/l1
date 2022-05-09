package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func first(jobs chan int) {
	for j := range jobs {
		fmt.Println(j)
		time.Sleep(time.Millisecond*50)
		//results <- j * 2
	}
	fmt.Println("Worker interrupted")

}

func second(quit chan bool) {
	for {
		select {
		case <-quit:
			fmt.Println("closing second")
			return
		default: 
			fmt.Println("Second goroutine running")
			time.Sleep(time.Millisecond*50)
		}
	}
}

func third(wg *sync.WaitGroup) {
	i:=1
	for i<10{
		fmt.Println("Third goroutine running ", i)
		i++
	}
	wg.Done()
}

func fourth(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("closing fourth")
			return
		default: 
			fmt.Println("4 goroutine running")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	quit := make(chan bool)
	jobs := make(chan int, 100)
	go first(jobs)
	jobs <- 1
	time.Sleep(time.Second)
	close(jobs)//первый способ - закрытие канала
	go second(quit)
	time.Sleep(time.Second)
	quit <- true//второй способ - отправить значение в канал
	wg.Add(1)
	go third(&wg)//третий способ - с помощью waitgroup.done
	wg.Wait()
	fmt.Println("Third goroutine interrupted")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)//четвертый способ - с помощью контекста
	defer cancel()
	go fourth(ctx)
	time.Sleep(time.Second*6)
	


}