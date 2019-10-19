package main

import "testing"

var input string = "omama"

func TestCountVowelsAndConsonants(t *testing.T) {
	output := map[string]int{"vowels":0, "consonants":0}
	expected := map[string]int{"vowels":2, "consonants":2}

	output = CountVowelsAndConsonants(input)
	
	t.Logf("Word : %s\n", input)
	t.Logf("Output : %v\n", output)
	t.Logf("Expected : %v\n", expected)

	if (output["vowels"] != expected["vowels"] || 
			output["consonants"] != expected["consonants"]) {
		t.Errorf("Error! expected (%v) got (%v)", expected, output)
	}

}

func BenchmarkCountVowelsAndConsonants(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountVowelsAndConsonants(input)
	}
}