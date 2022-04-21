package main

import "fmt"

func main() {
	var slice []int
	var a int

	swapper := true
	for swapper {
		fmt.Println("введите число")
		if n, err := fmt.Scanln(&a); err != nil || n != 1 {
			break
		}
		slice = append(slice, a)
	}

	fmt.Println(slice)
	InsertionSort(slice)
	fmt.Println(slice)
}

func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		x := slice[i]
		j := i
		for ; j >= 1 && slice[j-1] > x; j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = x
	}
}
