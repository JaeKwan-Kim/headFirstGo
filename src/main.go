package main

import "fmt"

func paintNeeded(width float64, height float64, be_print bool) float64 {
	area := width * height

	if be_print == true {
		fmt.Printf("\t %.2f liters needed\n", area/10.0)
	}

	return area / 10.0
}

func main() {
	var total float64

	total += paintNeeded(4.2, 3.0, true)
	total += paintNeeded(5.2, 3.5, true)

	fmt.Printf("Total : %0.2f liters\n", total)

}
