package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var a int64
	var pos int
	var bit byte
	fmt.Print("Введите натуральное число: ")
	_, err:=fmt.Scan(&a)
	if (err != nil) || (a < 1) {
		fmt.Println("Необходимо ввести натуральное число!", err)
	} else{
		s = strconv.FormatInt(a, 2)
		fmt.Print("Введите позицию бита: ")
		_, err:=fmt.Scan(&pos)
		if (err != nil) || (pos < 1) {
			fmt.Println("Необходимо ввести натуральное число!", err)
		} else {
			fmt.Print("Введите 0 или 1: ")
			_, err:=fmt.Scan(&bit)
			if (err != nil)   {
				fmt.Println("Необходимо ввести 0 или 1!", err)
			} else {
				str:=[]byte(s)
				if bit == 1 {
					str[len(s)-pos] = '1'
				} else {
					str[len(s)-pos] = '0'
				}
				
				//fmt.Println(pos, " ", bit, " ", string(str))
				s=string(str)
				fmt.Println(s)
				z,_:=strconv.ParseInt(s, 2, 64)
				fmt.Println(z)
			}
		}
	}
	
}