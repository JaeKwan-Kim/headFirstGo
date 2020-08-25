package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64, debugPrint bool) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a whdth of %2.f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %2.f is invalid", height)
	}

	area := width * height

	if debugPrint == true {
		log.Printf("%.2f liters needed", area/10.0)
	}

	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(4.2, 3.0, false)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Total : %0.2f liters\n", amount)
	}
}
