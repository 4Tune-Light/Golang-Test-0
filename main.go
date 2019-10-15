package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	var alive int
	var dead int
	var consonant bool

	vowels := []int32{'a', 'i', 'u', 'e', 'o'}
	check := []bool{false, false, false, false, false}

	fmt.Print("Masukkan Kata: ")
	fmt.Scanln(&word)

	for _, char := range strings.ToLower(word) {
		consonant = true

		for i, vowel := range vowels {
			if char == vowel {
				consonant = false
				if check[i] == false {
					alive++
					check[i] = true
				}
			}
		}

		if consonant == true {
			dead++
		}
	}

	fmt.Printf("Huruf Hidup: %d, Huruf Mati: %d", alive, dead)
}