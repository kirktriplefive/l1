package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="snow dog golang sun"
	fmt.Println(s)
	v:=strings.Split(s, " ")//берем слайс из слов строки 
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
        v[i], v[j] = v[j], v[i]//меняем элементы местами
    }
	s=strings.Join(v," ")//снова объединяем в строку
	fmt.Println(s)

}