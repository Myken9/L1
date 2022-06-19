package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	fmt.Println(Big("314159265358979323846264338327950288419716939937510", "14159265358979323846264338327950288419716939937510"))
}

func Big(a, b string) (answer string, err error) {
	na, ok := big.NewInt(0).SetString(a, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", a)
	}

	nb, ok := big.NewInt(0).SetString(b, 10)
	if !ok {
		return "", fmt.Errorf("ошибка преобразования строки в число: %s", b)
	}

	add := na.Add(na, nb).String()
	sub := na.Sub(na, nb).String()
	mul := na.Mul(na, nb).String()
	div := na.Div(na, nb).String()
	answer = "Addition = " + add + ", subtraction = " + sub + ", multiplication = " + mul + " division = " + div
	return answer, nil
}
