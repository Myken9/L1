package main

import (
	"os"
	"strconv"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов
//чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	Squares([]int{1, 2, 3, 4, 5})
}

func Squares(n []int) {
	var wg sync.WaitGroup

	wg.Add(len(n))

	for _, v := range n {
		go func(i int) {
			os.Stdout.WriteString(strconv.Itoa(i*i) + "\n")
			wg.Done()
		}(v)
	}
	wg.Wait()
}
