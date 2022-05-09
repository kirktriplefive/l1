package main

import "fmt"
func main() {
    v := "hello world"
    fmt.Println(typeof(v))
	a := 4
    fmt.Println(typeof(a))
	ch:=make(chan int)
	fmt.Println(typeof(ch))

}
func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)//выводим тип переменной
}