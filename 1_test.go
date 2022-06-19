package L1

import (
	"fmt"
	"testing"
)

//Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание
//методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayInfo() {
	fmt.Println("My name is ", h.Name, ". I am ", h.Age, " years old")
}

type Action struct {
	Human
}

func Test_test(t *testing.T) {
	h := Human{Name: "Bob", Age: 32}
	a := Action{h}
	a.SayInfo()
}
