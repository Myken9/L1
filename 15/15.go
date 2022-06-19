package main

import (
	"fmt"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/*
1)Глобальные переменные - зло
2) Символы в разных языках могут занимать разное колличесто байтов
*/
func main() {
	fmt.Println(someFunc())
}

func someFunc() string {
	v := createHugeString(1 << 10)

	runes := []rune(v)
	trueString := string(runes[:100])
	return trueString
}

func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "你"
	}

	return
}
