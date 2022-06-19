package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	row := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	newMap := Map{m: make(map[int][]float64)}

	for _, v := range row {
		newMap.store(int(v/10)*10, v)
	}
	fmt.Println(newMap)
}

type Map struct {
	m map[int][]float64
}

func (m *Map) load(key int) ([]float64, bool) {
	val, ok := m.m[key]
	return val, ok
}

func (m *Map) store(key int, value float64) {
	m.m[key] = append(m.m[key], value)
}
