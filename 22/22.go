package main

import (
	"fmt"
	"math/big"
)

func main() {
	f:=big.NewInt(36e17)//big.Int для крупных целых чисел, когда 18 квинтиллионов недостаточно;
	a:=big.NewInt(13e16)
	result:=new(big.Int)
	result=result.Div(f,a)
	fmt.Println("Деление", result)
	result=result.Add(f,a)
	fmt.Println("сложение", result)
	result=result.Sub(a,f)
	fmt.Println("Разность", result)
	result=result.Mul(f,a)
	fmt.Println("Перемножение", result)

}