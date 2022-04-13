package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("введите площадь круга")
	fmt.Scanln(&a)

	if a > 0 {
		var D = math.Sqrt(4 * a / math.Pi)
		var L = math.Pi * D
		fmt.Printf("%s %.2f\n", "Диаметр окружности равен:", D)
		fmt.Printf("%s %.2f\n", "Длина окружности равна:", L)
	} else {
		fmt.Println("Введите положительное число")
	}
}
