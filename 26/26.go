package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func main() {
	fmt.Println(unique("abcD"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}

func unique(s string) bool {
	s = strings.ToLower(s)
	a := strings.Split(s, "")
	entries := map[string]struct{}{}

	for _, v := range a {
		if _, ok := entries[v]; ok {
			return false
		}

		entries[v] = struct{}{}
	}

	return true
}
