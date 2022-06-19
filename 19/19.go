package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {
	fmt.Println(reverseString("Главрыба"))
	fmt.Println(reverseString("Information"))
	fmt.Println(reverseString("科技資訊"))
}

func reverseString(s string) (res string) {
	letters := strings.Split(s, "")

	for i := len(letters) - 1; i > -1; i-- {
		res += letters[i]
	}
	return
}
