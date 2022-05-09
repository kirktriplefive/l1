package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int//инкапсулированные параметры
	y int
}

func NewPoint(x,y int) *Point {
    return &Point{//конструктор
		x: x,
		y: y,
    }
}

func main() {
	var i,a int
	fmt.Print("Введите координату первой точки элемента для удаления ")
	_, err:=fmt.Scan(&i)
	if (err != nil) || (i < 1) {
		fmt.Println("Необходимо ввести натуральное число!", err)
	} else {
		fmt.Print("Введите номер элемента для удаления ")
		_, err=fmt.Scan(&a)
		if (err != nil) || (a < 1) {
			fmt.Println("Необходимо ввести натуральное число!", err)
		} else {
			NewPoint:=NewPoint(i, a)
			lenght:=math.Abs(float64(NewPoint.x-NewPoint.y))
			fmt.Println(lenght)
	
			
		}

		
	}
}

