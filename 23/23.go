package main

import "fmt"

func main() {
	slice:=[]int{1,2,3,4,5,6,7}
	var i int
	fmt.Print("Введите номер элемента для удаления ")
	_, err:=fmt.Scan(&i)
	if (err != nil) || (i < 1) {
		fmt.Println("Необходимо ввести натуральное число!", err)
	} else {
		i-=1
		copy(slice[i:], slice[i+1:])//выполняем сдвиг
		slice = slice[:len(slice)-1]//урезаем слайс на 1
		fmt.Println(slice)
	}

}