package assert

import "testing"

type TTestType interface {
	int | byte | string
}

func Equal[TValue TTestType](t *testing.T, expected TValue, actual TValue) {
	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}

func EqualMany[TArray TTestType](t *testing.T, expected []TArray, actual []TArray) {
	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}
