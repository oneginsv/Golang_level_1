package main

import "fmt"

func main() {
	var a int
	fmt.Println("введите a")
	fmt.Scanln(&a)
	x := []int{2}

	for i := 3; i < a; i += 2 {
		simple := true

		for _, j := range x {
			if i%j == 0 {
				simple = false
				break
			}
		}

		if simple {
			x = append(x, i)
		}
	}

	fmt.Println(x)
}
