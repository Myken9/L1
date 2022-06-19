package main

import (
	"os"
	"strconv"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
//результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	pipeNumbers([]int{1, 2, 3, 4, 5})
}

func pipeNumbers(num []int) {
	in := make(chan int)
	out := make(chan int)

	go func(out chan<- int) {
		for _, v := range num {
			in <- v
		}
		close(out)
	}(in)

	go func(in <-chan int, out chan<- int) {
		for v := range in {
			out <- v * 2
		}
		close(out)
	}(in, out)

	for v := range out {
		os.Stdout.WriteString(strconv.Itoa(v) + "\n")
	}
}
