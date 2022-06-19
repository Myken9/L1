package main

import "fmt"

//Поменять местами два числа без создания временной переменной.

func main() {
	a, b := 10, 57
	fmt.Println(a, b)
	changeValue(&a, &b)
	fmt.Println(a, b)
}

func changeValue(a, b *int) {
	*a, *b = *b, *a
}
