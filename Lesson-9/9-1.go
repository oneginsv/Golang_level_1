package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	config2 := parser.p_json()
	fmt.Println(&config2)

}
