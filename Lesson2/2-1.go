package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	fmt.Println("введите a")
	fmt.Scanln(&a)
	fmt.Println("введите b")
	fmt.Scanln(&b)

	if a > 0 && b > 0 {
		fmt.Printf("%s %.2f", "Площадь прямоугольника равна:", a*b)
	} else {
		fmt.Println("Вы ввели не число")
	}
}
