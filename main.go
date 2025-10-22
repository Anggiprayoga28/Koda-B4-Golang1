package main

import (
	"fmt"
	"go-practice/circle"
)

func main() {
	var radius float64
	var choice int

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("       KALKULATOR LINGKARAN")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	circle := circle.NewCircle(radius)

	if !circle.IsValid() {
		fmt.Println("Error: Jari-jari harus lebih besar dari 0")
		return
	}

	fmt.Println("\nPilih perhitungan:")
	fmt.Println("1. Hitung Luas saja")
	fmt.Println("2. Hitung Keliling saja")
	fmt.Println("3. Hitung Luas dan Keliling")
	fmt.Print("Pilihan Anda (1/2/3): ")
	fmt.Scan(&choice)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("               HASIL")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	switch choice {
	case 1:
		wide := circle.CalculateArea()
		fmt.Printf("Luas Lingkaran: %.2f\n", wide)
	case 2:
		around := circle.CalculateCircumference()
		fmt.Printf("Keliling Lingkaran: %.2f\n", around)
	case 3:
		wide := circle.CalculateArea()
		around := circle.CalculateCircumference()
		fmt.Printf("Luas Lingkaran: %.2f\n", wide)
		fmt.Printf("Keliling Lingkaran: %.2f\n", around)
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

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
