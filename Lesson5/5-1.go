package main

import "fmt"

func fibbo(n uint64) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbo(n-1) + fibbo(n-2)
}
func main() {
	var a uint64
	fmt.Println("введите порядковый номер числа Фибоначчи")
	fmt.Scanln(&a)
	fmt.Println("число Фибоначчи:", fibbo(a))
}
