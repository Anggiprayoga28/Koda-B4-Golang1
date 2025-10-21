package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	var choice int

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("       KALKULATOR LINGKARAN")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	if radius <= 0 {
		fmt.Println("Error")
		return
	}

	fmt.Println("\nPilih perhitungan:")
	fmt.Println("1. Hitung Luas saja")
	fmt.Println("2. Hitung Keliling saja")
	fmt.Println("3. Hitung Luas dan Keliling")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&choice)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("               HASIL")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	switch choice {
	case 1:
		luas := wide(radius)
		fmt.Printf("Luas Lingkaran: %.2f\n", luas)
	case 2:
		keliling := around(radius)
		fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
	case 3:
		luas, keliling := wideAround(radius)
		fmt.Printf("Luas Lingkaran: %.2f\n", luas)
		fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	var lagi string
	fmt.Print("\nHitung lagi? (y/n): ")
	fmt.Scan(&lagi)

	if lagi == "y" || lagi == "Y" {
		fmt.Println()
		main()
	} else {
		fmt.Println("\nTerima kasih")
	}
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
