package main

import (
	"fmt"
)

var pow = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

type MyMap struct {
	m map[int][]float64
}

func NewMap() *MyMap {
	return &MyMap{
		m: make(map[int][]float64),
	}
}

func (c *MyMap) Load(key int) ([]float64, bool) {
	val, ok := c.m[key]
	return val, ok
}

func (c *MyMap) Store(key int, value float64) {
	c.m[key] = append(c.m[key], value)
}

func main() {
	fmt.Println("Дана последовательность температурных колебаний:", pow)
	myMap := NewMap()
	for _, v := range pow {
		myMap.Store(int(v/10)*10, v)//загружаем значение в мап, где ключ - целое число от деления на 10, умноженное на 10, а значение - температура
	}
	fmt.Println(myMap)

}
