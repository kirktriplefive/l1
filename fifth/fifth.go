package main

import (
	"fmt"
	"time"
)

func receive(jobs chan int) {//функция для отправки значений в канал
	i := 0
	for {
		jobs <- i
		i++
	}
}


func worker(jobs chan int) {//функция для чтения из канала
	for j := range jobs {//пока в канал приходят значения функция работает
		fmt.Println(j)
		time.Sleep(time.Millisecond*50)
		//results <- j * 2
	}
	fmt.Println("Worker interrupted")

}

func main() {
	var a int
	
	fmt.Print("Введите количество секунд(натуральное число): ")
	_, err:=fmt.Scan(&a)
	if (err != nil) || (a < 1) {
		fmt.Println("Необходимо ввести натуральное число!", err)
	} else {
		jobs := make(chan int)
		go receive(jobs)
		go worker(jobs)
	}
	time.Sleep(time.Second*time.Duration(a))

}
