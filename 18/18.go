package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

func main() {
	counter := &counter{i: 0}
	counter.worker(5000)

}

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

func (c *counter) worker(iter int) {

	for i := 1; i <= iter; i++ {
		c.Add(1)
		go func(i int) {
			defer c.Done()
			fmt.Println(i)
			c.increment()
		}(i)
	}
	c.Wait()
	fmt.Println("count " + strconv.Itoa(c.i))
}
