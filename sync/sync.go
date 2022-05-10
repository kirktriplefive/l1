package main

import "fmt"

func myFunc(ch chan int, a int){
	a=a*a
	fmt.Println(a)
	ch<-a
}



func main() {
	a:=912
	ch:=make(chan int)
	go myFunc(ch, a)
	fmt.Println("hello")
	<-ch
	slice := make([]string,2, 2)
	func(slice []string) {
		slice[0] = "b"
		slice[1] = "b"
		slice = append(slice, "a")
		fmt.Print(slice)
	 }(slice)
	 fmt.Print(slice)
   
}