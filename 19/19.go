package main

import "fmt"

func revert(v []rune) string {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]//меняем местами 
	}
	return string(v)
}


func main() {
	s:="главрыба"
	v:=[]rune(s)//rune - int32(так как мы работаем с unicode)
	fmt.Println(v)
    fmt.Println(revert(v))
}