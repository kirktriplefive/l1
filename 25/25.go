package main

import (
    "fmt"
    "time"
)

var c chan int


func mySleep(duration time.Duration) {
	select {
    case <-time.After(duration)://с помощью тайм after ожидаем нужное время после чего 
    }
}

func main() {
	var i int
	var duration string
	fmt.Println("Введите время в формате 300ms, -1.5h или 2h45m")
	fmt.Scan(&duration)
	if dur, err := time.ParseDuration(duration); err!=nil {
		fmt.Println("Вы ввели время в неверном формате!")
	} else {
		for i=1; i<10; i++ {
			fmt.Println(i)
			mySleep(dur)
		}
	}
	
    
}