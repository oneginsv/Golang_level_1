package main

import "fmt"

func main() {
	var slice []int  // объявление слайса интов
	var a int  // кол-во повторений
	swapper := true
	// цикл ввода значений до тех пор, пока не введется НЕ число
	for swapper {
		fmt.Println("Введите число. Для выхода введите букву/ы")
		if n, err := fmt.Scanln(&a); err != nil || n != 1 {
			break
		}
		slice = append(slice, a)
	}
	Bubblesort(slice)
	fmt.Println(slice)
}

func  Bubblesort([] int slice)([] int result){
	for i:=0;i<len(slice);i++{
	   for j:=0;j<len(slice)-i-1;j++{
									// Сравните два числа
				  if slice[j]<slice[j+1]{
				  // Меняем местами два числа
							   slice[j],slice[j+1]=slice[j+1],slice[j]
				 }
	   }
  }
   return 
 }
