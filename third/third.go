package main

import (
	"fmt"
	"time"
)

func sum(c chan int){//создаем функцию для вычисления суммы квадратов
	var s int
	for {//функция работает пока приходят данные в канал
        val, ok := <-c
        if ok == false {
            break // если данные перестали приходить мы заканчиваем работу цикла и выводим сумму
        } else {
            s=s+val;
        }
    }
	fmt.Println(s)

}


func main() {
	numbers:= [5]int{2,4,6,8,10}
	start := time.Now()
	c := make(chan int)

	go sum(c)

	go func() {
		for _, n := range numbers {
			c<-n*n//отправляем квадрат элемента в канал
		}
		close(c)//закрываем канал
	}()

	elapsedTime := time.Since(start)
	fmt.Println("Total Time For Execution: " + elapsedTime.String())
	time.Sleep(time.Second)
}
