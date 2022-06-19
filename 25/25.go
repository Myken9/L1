package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	t := time.Now()
	fmt.Println("Start!")
	sleep(1 * time.Second)
	fmt.Printf("passed - %v\n", time.Since(t))
}

func sleep(t time.Duration) {
	<-time.After(t)
}
