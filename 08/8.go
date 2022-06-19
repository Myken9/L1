package main

import (
	"fmt"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	fmt.Printf("%d\n", changeBit(0, 63))
}
func changeBit(n int64, i int) (res int64) {
	return n ^ (1 << (i - 1))
}
