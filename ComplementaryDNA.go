/*

Codewars

Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". You have function with one side of the DNA (string, except for Haskell); you need to get the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

Example: (input --> output)

"ATTGC" --> "TAACG"
"GTAT" --> "CATA"
dnaStrand []        `shouldBe` []
dnaStrand [A,T,G,C] `shouldBe` [T,A,C,G]
dnaStrand [G,T,A,T] `shouldBe` [C,A,T,A]
dnaStrand [A,A,A,A] `shouldBe` [T,T,T,T]
*/

package main

import (
	"fmt"
)

func DNAStrand(dna string) string {

	A := rune('A')
	T := rune('T')
	C := rune('C')
	G := rune('G')

	/*
		it is necessary to know that in golang,we cannot
		convert a string to a single rune, we can only convert a single
		char to a single rune, but we can convert strings to a slice of runes
	*/

	stringInRune := []rune(dna)

	for i := 0; i < len(stringInRune); i++ {

		if string(stringInRune[i]) == "A" {
			stringInRune[i] = T
		} else if string(stringInRune[i]) == "T" {
			stringInRune[i] = A
		}
		if string(stringInRune[i]) == "C" {
			stringInRune[i] = G
		} else if string(stringInRune[i]) == "G" {
			stringInRune[i] = C
		}
	}
	return string(stringInRune)

}

func main() {

	DNA := "ATTGC"
	fmt.Println(DNAStrand(DNA))

}

/*

some other persons answer, using the string package

*/

/*import (
    "strings"
)

func DNAStrand(dna string) string {
  // your code here
  replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
  return(replacer.Replace(dna))
}*/
