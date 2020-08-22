package main

import (
	"errors"
	"fmt"
)

func paintNeeded(width float64, height float64, be_print bool) (error, float64) {
	area := width * height

	if area < 0 || width < 0 {
		err := errors.New("area OR width Can't")
		return err, 0
	}

	if be_print == true {
		fmt.Printf("\t %.2f liters needed\n", area/10.0)
	}

	return nil, area / 10.0
}

func main() {
	var total float64

	_, total = paintNeeded(4.2, 3.0, true)
	_, total = paintNeeded(5.2, 3.5, true)

	fmt.Printf("Total : %0.2f liters\n", total)
	fmt.Printf("Total : %0.2f liters\n", total)
}
