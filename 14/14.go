package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel
//из переменной типа interface{}.

func main() {
	fmt.Println(Type1(23.354))
	fmt.Println(Type2(23.354))
}

func Type1(object interface{}) string {
	switch object.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case struct{}:
		return "struct"
	case chan interface{}:
		return "channel"
	case interface{}:
		return "interface"
	case nil:
		return "nil"
	default:
		return "unknown"
	}
}

func Type2(object interface{}) reflect.Type {
	return reflect.TypeOf(object)
}
