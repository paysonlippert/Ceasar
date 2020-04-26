package main

import "fmt"

func main() {
	offset := 20
	result, _ := ceasar("The Quick Brown Fox Jumped Over the Moon", offset)
	fmt.Println(result)
	result, _ = ceasar(result, -offset)
	fmt.Println(result)
}

func ceasar(plainText string, offset int) (string, error) {
	resultRunes := []rune{}
	textRune := []rune(plainText)
	offsetRune := rune(offset)
	for _, rune := range textRune {
		newRune := rune + offsetRune
		resultRunes = append(resultRunes, newRune)
	}
	return string(resultRunes), nil
}
