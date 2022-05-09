package main

import "fmt"

func main() {
	n := 0
	fmt.Println(&n)
	if true {
	   n := 1
	   n++
	   fmt.Println(&n)
	   fmt.Println(n)
	}
	fmt.Println(&n)
	fmt.Println(n)
  }

  