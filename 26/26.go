package main

import (
	"fmt"
	"strings"
)

func checking(s []rune) bool {
	m:=make(map[rune]int)
	for i:=0; i<len(s); i++ {
		m[s[i]]+=1//мапим наш слайс рун по ключу - значению и значение - счетчик повторения
	}
	for _, value := range m {
		if value>1 {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Println("Строку, которую хотите проверить ")
	fmt.Scan(&str)
	str=strings.ToLower(str)
	s:=[]rune(str)
	f:=checking(s)
	fmt.Println(f)

}