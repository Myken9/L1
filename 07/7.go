package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.

func main() {
	a := syncmap{m: make(map[int]int)}

	for i := 0; i < 10; i++ {
		go func(i int) {
			a.add(1, i)
		}(i)
	}

	time.Sleep(1 * time.Second)

	for k, v := range a.m {
		fmt.Printf("%v %v\n", k, v)
	}

}

type syncmap struct {
	sync.Mutex
	m map[int]int
}

func (s *syncmap) add(key, value int) {
	s.Lock()
	s.m[key] = value
	s.Unlock()
}

func (s *syncmap) get(key int) int {
	s.Lock()
	defer s.Unlock()

	return s.m[key] // zero value if not found
}
