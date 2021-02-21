# Cook It 

A take on FizzBuzz, cookit is a go library that evaluates a list of numbers.   
If a number in the list is divisible by 3, it will print Crackle instead of the number. If it's divisible by 5, it will print Pop. If it's divisible by both 3 and 5, it will print CracklePop.

## Prerequisites

- go 1.15 

## Running Tests 

- `go test github.com/kimschles/cookit`

## Usage 

```
package main

import (
	cp "github.com/kimschles/cookit"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}

	cookedSlice := cp.CookIt(numbers)

	cp.ShowIt(cookedSlice)

}
```