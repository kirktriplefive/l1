package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	ar := []int{5, 3, 2, 1, 7, 0, 8, 12}
	start := time.Now()
	sort.Ints(ar)
	elapsedTime := time.Since(start)
	fmt.Println(ar, elapsedTime)
	start = time.Now()
	Quicksort(ar)
	elapsedTime = time.Since(start)
	fmt.Println(ar, elapsedTime)
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		tmp := ar[left]
		ar[left] = ar[right]
		ar[right] = tmp
	}
}
