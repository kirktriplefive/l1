package main

import (
	"fmt"
	"sort"
)

func main() {
	ar := []int{0,3,1,5,6,11,23,2,21}
	f:=SearchInIntSlice(ar, 5)
	fmt.Println(f)
}

func SearchInIntSlice(ar []int, s int) (result bool) {
    sort.Ints(ar) 
	fmt.Println(ar)
    lowKey := 0 // первый индекс
    highKey := len(ar) - 1 // последний индекс
    if (ar[lowKey] > s) || (ar[highKey] < s) {
        return // нужное значение не в диапазоне данных
    }
    for lowKey <= highKey { 
        // уменьшаем список рекурсивно
        mid := (lowKey + highKey) / 2 // середина
        if ar[mid] == s {
            result = true // мы нашли значение
            return
        } else if ar[mid] < s { 
            // если поиск больше середины - мы берем только блок с большими значениями увеличивая lowKey
            lowKey = mid + 1
            continue
        } else if ar[mid] > s { 
            // если поиск меньше середины - мы берем блок с меньшими значениями уменьшая highKey
            highKey = mid - 1
        }
    }
    return
}