package main

import (
	"fmt"
	"time"
)

func f(i int) {
    fmt.Println(i*i, time.Now())//функция для вывода квадратов
}


func main() {
	numbers:= [5]int{2,4,6,8,10}//массив
	start := time.Now()

	for _, n := range numbers {
		go f(n)//вызов функции в горутине для конкурентного расчета квадратов 
	}
  
	elapsedTime := time.Since(start)
  
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
	time.Sleep(time.Second)//ожидание завершения работы всех горутин(можно реализовать с помощью WaitGroup, что будет показано в других заданиях)
}