package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type counter struct {
	i int
	sync.Mutex
	sync.WaitGroup
}

func (c *counter) increment() {
	c.Lock()
	c.i++
	c.Unlock()
}

func worker(c *counter, iter int) {

	for i := 1; i <= iter; i++ {
		c.Add(1)
		go func(i int) {
			defer c.Done()
			fmt.Println(i)
			c.increment()
		}(i)
	}
	c.Wait()

}

func main() {
	counter := &counter{i: 0}
	worker(counter, 500)
	fmt.Println("count " + strconv.Itoa(counter.i))
}
