package main

import (
	"flag"
	"log"
)

var (
	maxRows = flag.Int("max-rows", 7, "Максимальное число строк,которые нужно прочитать")
	columns = flag.String("columns", "", "Разделённый запятыми список столбцов,которые нужно вывести")
)

func main() {
	flag.Parse() // Сообщаем библиотеке flag, что необходимо считать флаги
	// Имя файла передаётся не через флаги, а как аргументы командной строки flag.Arg.
	if flag.NArg() != 1 {
		log.Fatal("Неверно задано количество аргументов командной строки.Проверьте, что вы задали имя файла.")
	}

}