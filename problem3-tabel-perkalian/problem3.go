package main

import "fmt"

func CetakTabelPerkalian(number int) {
	// your code here
	for i := 1; i <= number; i++ { // looping baris tabel perkalian
		for j := 1;  j <= number; j++ { // looping baris tabel nomor kolom 
			fmt.Printf("%4d", i*j) // mencetak miliki panjang 4 karakter 

		}
		fmt.Println()
	}
}
func main() {
	CetakTabelPerkalian(9)
}
