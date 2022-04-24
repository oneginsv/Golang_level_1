package main

import "fmt"

var myMap map[uint64]uint64

func fibbo(n uint64) uint64 {
	if n < 2 {
		return myMap[n]
	}
	var ok bool
	var n2 uint64
	if n2, ok = myMap[n-2]; !ok {
		n2 = fibbo(n - 2)
	}
	var n1 uint64
	if n1, ok = myMap[n-1]; !ok {
		n1 = fibbo(n - 1)
	}
	myMap[n] = n1 + n2
	return myMap[n]
}

func main() {
	var a uint64
	fmt.Println("введите порядковый номер числа Фибоначчи")
	fmt.Scanln(&a)
	myMap = make(map[uint64]uint64)
	myMap[0] = 0
	myMap[1] = 1
	fmt.Println("число Фибоначчи:", fibbo(a))
}
