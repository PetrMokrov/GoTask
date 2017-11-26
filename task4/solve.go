package main

import (
	"fmt"
	"unicode"
)

func RemoveEven(InputSlice []int) (OutputSlice []int) {
	for _, val := range InputSlice {
		if val%2 != 0 {
			OutputSlice = append(OutputSlice, val)
		}
	}
	return
}

func PowerGenerator(Power int) func() int {
	currentPower := 1
	initialPower := Power
	return func() int {
		currentPower = currentPower * initialPower
		return currentPower
	}
}

func DifferentWordsCount(Str string) int {
	usingStrings := make(map[string]bool)
	var currentStr string
	for _, valRune := range Str {
		if unicode.IsLetter(valRune) {
			currentStr = currentStr + string(unicode.ToLower(valRune))
		} else {
			_, ok := usingStrings[currentStr]
			if !ok && currentStr != "" {
				usingStrings[currentStr] = true
			}
			currentStr = ""
		}
	}
	//fmt.Println(usingStrings)
	return len(usingStrings)
}
/*
func main() {
	input := []int{0, 3, 2, 5}
	result := RemoveEven(input)
	fmt.Println(result) // Должно напечататься [3 5]
	gen := PowerGenerator(2)
	fmt.Println(gen()) // Должно напечатать 2
	fmt.Println(gen()) // Должно напечатать 4
	fmt.Println(gen()) // Должно напечатать 8
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
	// Должно напечатать 2
}
*/
