package cookit

import (
	"fmt"
	"strconv"
)

// CookIt evaluates a slice of numbers. If the number is divisible by 3, it adds "Crackle" to a new slice, if the number is divisible by 5, it adds "Pop", and if it's divisible by both 3 and 5, it adds "CracklePop"
func CookIt(numbers []int) []string {
	var alteredSlice []string
	for _, num := range numbers {
		if num%3 == 0 && num%5 == 0 {
			alteredSlice = append(alteredSlice, "CracklePop")
		} else if num%3 == 0 {
			alteredSlice = append(alteredSlice, "Crackle")
		} else if num%5 == 0 {
			alteredSlice = append(alteredSlice, "Pop")
		} else {
			alteredSlice = append(alteredSlice, strconv.Itoa(num))
		}
	}
	return alteredSlice
}

// ShowIt prints all strings in a slice
func ShowIt(cookedSlice []string) {
	for _, item := range cookedSlice {
		fmt.Println(item)
	}
}
