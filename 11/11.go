package main

import "fmt"

func main() {
	first := []int{10, 11, 2, 45, 6, 7}
	second := []int{12, 3, 45, 7, 1, 9, 123, 4, 10}
	result := []int{}
	m := make(map[int]int)//инициализируем мап
	for _, v := range first {
		m[v] += 1//в мап по ключу - элементу массива увеличиваем счетчик
	}
	for _, val := range second {
		m[val] += 1
	}
	for key, value := range m {
		if value > 1 {
			result = append(result, key)//если счетсик >1, это значит что элемент есть в обоих массивах.
		}
	}
	fmt.Println(result)
}
