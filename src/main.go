package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("head first go"))
	fmt.Println("Hello, \nGo!")
	fmt.Println("Hello, \tGo!")
	fmt.Println("Quotes:\"\"")
	fmt.Println("Backslash:\\")

	// rune 룬
	fmt.Println('A')
	fmt.Println('B')
	fmt.Println('\t')

	//boolean
	fmt.Println(true)
	fmt.Println(false)
}
