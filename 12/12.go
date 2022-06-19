package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range strings {
		set[v] = struct{}{}
	}

	fmt.Println(set)
}
