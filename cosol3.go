package main

import "fmt"

func main() {
	var beratBadan, tinggiBadan, bmi float64
	fmt.Print("Masukan Berat Badan (kg): ")
	fmt.Scan(&beratBadan)
	fmt.Print("Masukan Tinggi Badan (m): ")
	fmt.Scan(&tinggiBadan)
	bmi = beratBadan / (tinggiBadan * tinggiBadan)
	fmt.Printf("BMI anda: %.2f", bmi)
}
