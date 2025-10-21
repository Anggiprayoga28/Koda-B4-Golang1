package main

import (
	"fmt"
	"math"
)

func main() {
	circleArea := wide(24)
	circumference := around(24)
	fmt.Println(circleArea, circumference)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	wide, around := wideAround(24)
	fmt.Print("Wide: ", wide, " Around: ", around)
}

func wide(rad float64) float64 {
	total := math.Pi * math.Pow(rad, 2)
	return total
}

func around(rad float64) float64 {
	total := 2 * math.Pi * rad
	return total
}

func wideAround(rad float64) (float64, float64) {
	totalArea := math.Pi * math.Pow(rad, 2)
	totalCircumference := 2 * math.Pi * rad
	return totalArea, totalCircumference
}
