package L1

import (
	"testing"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….)
//с использованием конкурентных вычислений.

func SumSquares(n []int) (sum int) {
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

func Test_SumSquares(t *testing.T) {
	if s := []int{2, 4, 6, 8, 10}; SumSquares(s) != 220 {
		t.Error("Expected 220, got ", SumSquares(s))
	}
}
