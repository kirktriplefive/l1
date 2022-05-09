package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	fmt.Println(cap(slice))
	func(slice []string) {
	   slice = append(slice, "a")
	   fmt.Println(cap(slice))
	   slice[0] = "b"
	   slice[1] = "b"
	   fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
  }
  