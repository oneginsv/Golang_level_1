package main

import "fmt"

func main() {
	var slice []int // объявление слайса интов
	var a int       // кол-во повторений
	swapper := true
	// цикл ввода значений до тех пор, пока не введется НЕ число
	for swapper {
		fmt.Println("Введите число. Для выхода введите букву/ы")
		if n, err := fmt.Scanln(&a); err != nil || n != 1 {
			break
		}
		slice = append(slice, a)
	}
	Selectsort(slice)
	fmt.Println(slice)
}

func Selectsort(slice []int) (result []int) {
	len := len(slice)
	for i := 0; i < len; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if slice[minIndex] > slice[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			slice[minIndex], slice[i] = slice[i], slice[minIndex]
		}
	}
	result = slice
	return
}
