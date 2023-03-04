package main

import "fmt"

func DrawXYZ(n int) string {
	// your code here
	var result string
	for i := 1; i <= n; i++ { // melooping untuk baris pada pola
		for j := 1; j <= n; j++ { // melooping untuk indeks kolom 
            if (i+j)%3 == 0 {	// jika indeks baris dan kolom berjumlah kelipatan 3, maka huruf x ditambahkan ke variabel result
				result += "X " // menambahkan karakter X ke variabel  
			} else if (i+j)%2 == 1{ // jika hasil penjumlahan indeks dan kolom menghasilakn bilangan ganjil, maka huruf Y ditambahkan ke variabel result
				result += "Y "// menambahkan karakter Y ke variabel 
			} else {
				result += "Z " // menambahkan karakter Z ke variabel
            }
		}	
		result += "\n" // menambahkan karakter newline ke variabel result
	}	
	return result// mengembalikan pola huruf X, Y, Z disimpan divariabel risult			
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}