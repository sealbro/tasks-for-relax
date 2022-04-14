package assert

import (
	"testing"
)

type TTestType interface {
	int | uint | int32 | uint32 | byte | string | bool
}

func True(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Expected: true, Actual: %v.", actual)
	}
}

func False(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Expected: false, Actual: %v.", actual)
	}
}

func Equal[TValue TTestType](t *testing.T, expected TValue, actual TValue) {
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v.", expected, actual)
	}
}

func EqualMany[TArray TTestType](t *testing.T, expected []TArray, actual []TArray) {
	if len(expected) != len(actual) {
		t.Errorf("\nDifferent length!!!\nExpected: %v,\nActual: %v.", expected, actual)
		return
	}

	for i, expectedItem := range expected {
		actualItem := actual[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}

func EqualMultiDimensional[TArray TTestType](t *testing.T, expected [][]TArray, actual [][]TArray) {
	if len(expected) != len(actual) {
		t.Errorf("\nDifferent length!!!\nExpected: %v,\nActual: %v.", expected, actual)
		return
	}

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			actualItem := actual[i][j]
			expectedItem := expected[i][j]
			if actualItem != expectedItem {
				t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
				return
			}
		}
	}
}
