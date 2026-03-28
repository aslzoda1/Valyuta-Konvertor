package main

import (
	"fmt"
)

func main() {

	var summa float64
	var kurs float64
	var tanlov int

	fmt.Println("=== Go Valyuta Konvertori ===")
	fmt.Println("1. USD -> UZS (Dollardan Somga)")
	fmt.Println("2. UZS -> USD (Somdan Dollarga)")
	fmt.Print("Tanlovingizni kiriting (1 yoki 2): ")
	fmt.Scanln(&tanlov)

	if tanlov == 1 {
		fmt.Print("Dollar miqdorini kiriting: ")
		fmt.Scanln(&summa)
		fmt.Print("1 AQSH dollari necha so'm? (Kurs): ")
		fmt.Scanln(&kurs)

		natija := summa * kurs
		fmt.Printf("%.2f USD = %.2f UZS\n", summa, natija)

	} else if tanlov == 2 {
		fmt.Print("So'm miqdorini kiriting: ")
		fmt.Scanln(&summa)
		fmt.Print("1 AQSH dollari necha so'm? (Kurs): ")
		fmt.Scanln(&kurs)

		if kurs == 0 {
			fmt.Println("Xato: Kurs 0 bo'lishi mumkin emas!")
		} else {
			natija := summa / kurs
			fmt.Printf("%.2f UZS = %.2f USD\n", summa, natija)
		}

	} else {
		fmt.Println("Noto'g'ri tanlov kiritildi.")
	}

	fmt.Println("---------------------------")
	fmt.Println("Dastur tugadi. Rahmat!")
}
