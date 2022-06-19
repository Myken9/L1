package main

import "fmt"

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Println(sumSquares(s))
}

func sumSquares(n []int) (sum int) {
	intCh := make(chan int)
	square := func(d int, intCh chan<- int) {
		intCh <- d * d
	}

	for _, v := range n {
		go square(v, intCh)
	}

	for i := 0; i < len(n); i++ {
		sum += <-intCh
	}
	return
}
