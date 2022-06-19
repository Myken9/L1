package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	t := time.Now()
	fmt.Println("Start!")
	Sleep(1 * time.Second)
	fmt.Printf("passed - %v\n", time.Since(t))
}

func Sleep(t time.Duration) {
	<-time.After(t)
}
