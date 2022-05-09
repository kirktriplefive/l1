package main

import (
    "fmt"
    "unsafe"
)


func main() {

    s := struct{}{}

    fmt.Printf("struct{}{} size: %v\n", unsafe.Sizeof(s))
}