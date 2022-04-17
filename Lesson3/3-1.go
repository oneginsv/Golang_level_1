package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Println("введите a")
	fmt.Scanln(&a)
	fmt.Println("введите b")
	fmt.Scanln(&b)
	fmt.Println("введите арифмитическую операцию")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делим на ноль - ошибка!")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "!":
		res = 1
		for a != 0 {
			res = res * a
			a--
		}
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
