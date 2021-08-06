package main

import (
	"fmt"
	"strings"

	input "github.com/copy-ninja1/acronym/pkg"
)

func main() {
	// text :=  pkg("Enter phrase to get acronym :")
	text := input.GetInput("Enter phrase to get acronym :")
	fmt.Println("text : ", text)
	acro := GetAcronym(text)
	fmt.Println("The acronym for ",text," is",acro)

}

func SplitTextToWords(text string) []string {
	words := strings.Split(text, " ")
	return words
}

func GetAcronym(words string) string {
	splitedText := SplitTextToWords(words)

	txtInitCount := 0
	var acronym string

	for txtInitCount < len(splitedText) {
		acronym += strings.ToUpper(splitedText[txtInitCount][0:1])
		txtInitCount++
	}
	return acronym
}
