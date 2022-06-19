package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	fmt.Println(sayLikeYoda("snow dog sun"))
	fmt.Println(sayLikeYoda("эта программа перевораячивает слова"))
}

func sayLikeYoda(s string) (res string) {
	letters := strings.Split(s, " ")

	for i := len(letters) - 1; i > -1; i-- {
		res += letters[i] + " "
	}
	return
}
