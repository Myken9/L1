package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

func main() {
	fmt.Println(type1(23.354))
	fmt.Println(type2(23.354))
}

func type1(object interface{}) string {
	switch object.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "channel"
	default:
		return "unknown"
	}
}

func type2(object interface{}) reflect.Type {
	return reflect.TypeOf(object)
}
