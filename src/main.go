package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))

	//var quantity int
	//var length, width float64
	//var customerName string
	//
	//quantity = 4
	//customerName = "Damon Cole"
	//length, width = 1.2, 2.4

	var quantity = 4
	var customerName = "Damon Cole"
	var length, width = 1.2, 2.4

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

}
