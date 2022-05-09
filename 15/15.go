package main

import (
	"fmt"
	"strings"
)

func createHugeString(len int) (str string) {
  for i := 0; i < len; i++ {
    str += "а"
  }
  return str
}


func someFunc() {
    v := createHugeString(1<<10)
    fmt.Println(len(v))//но создавали мы строку из 1024 символов. Проблема в том, что у нас символ занимает 2 байта, а слайс
    //берет побайтово
    vc := v[:100]                     // берем слайс
    fmt.Println(vc, len(vc)) // печатаем результат (неверный()
 
 }

func someFuncNew() {
  v := strings.Split(createHugeString(1<<10), "")//слайс символов из строки
  vc := v[:100]                     // берем слайс
  s:=strings.Join(vc,"")
  fmt.Println(vc, len(vc), cap(vc)) // печатаем результат
  fmt.Println(s)

}



func main() {
  someFunc()
  someFuncNew()
}