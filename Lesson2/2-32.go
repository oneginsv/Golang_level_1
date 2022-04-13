package main

import (
	"fmt"
)

func main() {
	//Запрос числа (может быть введено не число)
	var str int
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&str)

	if str >= 100 && str <= 999 {
		fmt.Printf("%s %d\n", "Количество сотен равно:", str/100)
		fmt.Printf("%s %d\n", "Количество десятков равно:", (str%100)/10)
		fmt.Printf("%s %d\n", "Количество единиц равно:", str%10)
	} else {
		fmt.Println("Error")
	}
}

func sot() {

	values := []int{0, 1, 2, 3, 4, 5, 6}

	for _, x := range values {

		w := x

		w *= 2

		fmt.Println(w)
	}
}
