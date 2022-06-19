package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	slice := []int{-4, -1, 0, 0, 5, 7, 9, 12, 15, 15, 17, 20, 22}
	element, err := binarySearch(slice, 22)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(element)
	}

}

func binarySearch(slice []int, value int) (int, error) {
	err := errors.New("no such element " + strconv.Itoa(value))

	left, right := 0, len(slice)-1
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == value {
			return mid, nil
		}

		if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0, err
}
