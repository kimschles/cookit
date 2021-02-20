package cookit

import (
	"strconv"
)

//CookIt scans a list of numbers. If the number is divisible by 3, it prints Crackle, if the number is divisible by 5, it prints Pop, f it's divisible by both 3 and 5, print CracklePop
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
