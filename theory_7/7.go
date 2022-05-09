package main

import "fmt"

func main() {
	ints := map[int]int{
		1: 3711,
		2:  2138,
		7: 1908,
		9: 912,
	}
	ints[5]=34
	ints[4]=8
	ints[89]=67
	fmt.Println(ints)
	
}