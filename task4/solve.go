package main

import (
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

func DifferentWordsCount(Str string) (resCount int) {
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
	_, ok := usingStrings[currentStr]
	if !ok && currentStr != "" {
		usingStrings[currentStr] = true
	}
	resCount = len(usingStrings)
	return
}
