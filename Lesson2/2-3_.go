package main

import (
	"fmt"
	"strconv"
	"strings"
)

package main

import (
	"fmt"
)

func main() {
	//Запрос числа 
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



func main2() {
	//Запрос числа (может быть введено не число)
	var str2 string
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&str2)

	// объявление массива
	arr := strings.Split(str2, "")
	var ints []int
	for _, s := range arr {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		} else {
			fmt.Println("Введите трехзначное число")
		}
	}
	fmt.Printf("%s %d\n", "Количество сотен равно:", ints[0])
	fmt.Printf("%s %d\n", "Количество десятков равно:", ints[1])
	fmt.Printf("%s %d\n", "Количество единиц равно:", ints[2])
}
