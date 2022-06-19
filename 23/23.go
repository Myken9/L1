package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	a, err := deleteElement(a, 7)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func deleteElement(slice []int, i int) ([]int, error) {
	if i < 0 || i >= len(slice) {
		return nil, fmt.Errorf("index does not exist")
	}
	slice = append(slice[:i], slice[i+1:]...)
	return slice, nil
}
