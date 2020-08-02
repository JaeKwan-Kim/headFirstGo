package main

import (
	"fmt"
	"reflect"
)

func main() {
	myInt := 2
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))

	length := 1.2
	width := 2

	//length = float64(width)
	//fmt.Println(length)

	width = int(length)
	fmt.Println(width)
}
