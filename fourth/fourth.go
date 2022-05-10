package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func receive(ingest chan int){
	i := 0
    for { 
			ingest <-i
			i++
			time.Sleep(time.Millisecond * 100)
        
    }
}

func starting(ctx context.Context, jobs, ingest chan int) {
	for {
		select {
		case job := <-ingest:
			jobs <- job
		case <-ctx.Done():
			fmt.Println("closing jobs")
			close(jobs)
			fmt.Println("closed jobs")
			return
		}
	}
}

func worker(id int, wg *sync.WaitGroup, jobs chan int) {
	defer wg.Done()
	for j := range jobs {
        fmt.Println(id," ", j)
		time.Sleep(time.Second)
    }
	fmt.Printf("Worker %d interrupted\n", id)

}

func main() {
	var a int
	
	fmt.Print("Введите количество воркеров(натуральное число): ")//вводим кол-во воркеров
	_, err:=fmt.Scan(&a)
	if (err != nil) || (a < 1) {
		fmt.Println("Необходимо ввести натуральное число!", err)
	} else {
		var wg sync.WaitGroup//инициализируем переменную типа WaitGroup для синхронизации работы горутин
		sigs := make(chan os.Signal, 1)
		ingest := make(chan int, 1)
		jobs := make(chan int, 100)
		ctx, cancel := context.WithCancel(context.Background())//инициализируем новый контекст с функцией отмены
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)//вызываем функцию, которая запишет сигнал ^C в sigs	
		go receive(ingest)//запускаем горутину, которая будет отправлять данные в канал ingest
		go starting(ctx, jobs, ingest)//запускаем горутину, которая будет принимать значения из каналов и выполнять проверку на завершение
		for i := 0; i < a; i++ {
			wg.Add(1)
			go worker(i,&wg,jobs)//запускаем воркера
		}
		<-sigs//ожидаем сигнал завершения
		cancel()//заупскаем функцию отмены 
    	wg.Wait()
	}
}
