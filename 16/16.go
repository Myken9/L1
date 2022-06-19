package main

import (
	"fmt"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	a := []int{25, 5, 9, 6, 8, 7, 2, 10, 1, 20}
	Quicksort(a)
	fmt.Println(a)
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
	// Выбираем опорный элемент
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

		ar[left], ar[right] = ar[right], ar[left]
	}
}
