package main

import (
	"fmt"
	"strings"
)

func CountVowelsAndConsonants(word string) map[string]int{
	var consonant bool
	var alive int
	var dead int
	
	data := make(map[string]int)
	vowels := []int32{'a', 'i', 'u', 'e', 'o'}
	check := []bool{false, false, false, false, false}

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

	data["vowels"] = alive
	data["consonants"] = dead

	return data
}

func main() {
	var word string
	data := make(map[string]int)

	fmt.Print("Masukkan Kata: ")
	fmt.Scanln(&word)

	data = CountVowelsAndConsonants(word)

	fmt.Printf("Huruf Hidup: %d, Huruf Mati: %d", data["vowels"], data["consonants"])

}