package cookit

import (
	"reflect"
	"testing"
)

func TestCrackle(t *testing.T) {
	testNumbers := []int{33, 36, 99, 22}
	returnSlice := []string{"Crackle", "Crackle", "Crackle", "22"}
	output := CookIt(testNumbers)
	if reflect.DeepEqual(output, returnSlice) != true {
		t.Errorf("TestCrackle failed, got: %v, wanted %v", output, returnSlice)
	}

}

func TestPop(t *testing.T) {
	testNumbers := []int{5, 40, 100, 79}
	returnSlice := []string{"Pop", "Pop", "Pop", "79"}
	output := CookIt(testNumbers)
	if reflect.DeepEqual(output, returnSlice) != true {
		t.Errorf("TestPop failed, got: %v, wanted %v", output, returnSlice)
	}
}

func TestCracklePop(t *testing.T) {
	testNumbers := []int{75, 15, 60, 85}
	returnSlice := []string{"CracklePop", "CracklePop", "CracklePop", "Pop"}
	output := CookIt(testNumbers)
	if reflect.DeepEqual(output, returnSlice) != true {
		t.Errorf("TestCracklePop failed, got: %v, wanted %v", output, returnSlice)
	}
}

// TO DO: finish this test
func TestShowIt(t *testing.T) {
	testNumbers := []int{49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	cookedSlice := CookIt(testNumbers)
	ShowIt(cookedSlice)
}
