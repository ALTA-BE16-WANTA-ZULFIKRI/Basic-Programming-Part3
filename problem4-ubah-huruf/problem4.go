package main

import (
	"fmt"
	"strings"
)

func UbahHuruf(sentence string) string {
	// your code here
	var alphabetBob = "KLMNOPQRSTUVWXYZABCDEFGHIJ" // declare variabel named alphabetic
	var encryptedSentence string // declare encrypted variable

	for _, ch := range sentence { // mengiterasi setiap karakter pada sentence 
		if ch == ' ' {
			encryptedSentence += " " // jika spaci karakter tambahkan ke encrypted 
		} else {
			idx := strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", ch) // jika bukan karakter dienkrisi ke variabel encrypt
			encryptedSentence += string(alphabetBob[idx])
		}
	}
	return encryptedSentence // mengembalikan nilai dari variabel setelah dienkripsi 
}

func main() {
	fmt.Println(UbahHuruf("SEPULSA OKE"))     // COZEVCK YUO
	fmt.Println(UbahHuruf("ALTERRA ACADEMY")) // KVDOBBK KMKNOWI
	fmt.Println(UbahHuruf("INDONESIA"))       // SXNYXOCSK
	fmt.Println(UbahHuruf("GOLANG"))          // QYVKXQ
	fmt.Println(UbahHuruf("PROGRAMMER"))      // ZBYQBKWWOB
}
