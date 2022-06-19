package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	fmt.Println(intersection1([]int{1, 2, 6, 23, 95}, []int{3, 54, 6, 346, 35, 23, 95}))
	fmt.Println(intersection2([]int{1, 2, 6, 23, 95}, []int{3, 54, 6, 346, 35, 23, 95}))
}

//O(n)
func intersection1(slice1, slice2 []int) (result []int) {
	m := make(map[int]struct{})
	for _, v := range slice1 {
		m[v] = struct{}{}
	}

	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	return
}

//O(n^2)
func intersection2(slice1, slice2 []int) (result []int) {
	for _, v := range slice1 {
		for _, k := range slice2 {
			if v == k {
				result = append(result, v)
			}
		}
	}
	return
}
