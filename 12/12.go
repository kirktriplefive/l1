package main

import "fmt"

//подобно 11(записываем в мап по ключу - значению из множества строк и увеличиваем счетчик)

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)
	result := []string{}
	for _, s := range strings {
		m[s] += 1
	}
	for key := range m {
		result = append(result, key)
	}
	fmt.Println(result)
}
