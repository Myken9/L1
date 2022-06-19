package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

func main() {

	h := human{Name: "Bob", Age: 32}
	a := action{h}
	a.sayInfo()

}

type human struct {
	Name string
	Age  int
}

func (h *human) sayInfo() {
	fmt.Println("My name is ", h.Name, ". I am ", h.Age, " years old")
}

type action struct {
	human
}
