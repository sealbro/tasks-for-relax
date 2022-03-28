package rotate_array

import "testing"

func TestSearchCase1(t *testing.T) {
	actualArr := []int{1, 2, 3, 4, 5, 6, 7}
	target := 3
	expectedArr := []int{5, 6, 7, 1, 2, 3, 4}

	rotate(actualArr, target)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("\nExpected: %v,\nActual: %v.", expectedArr, actualArr)
			break
		}
	}
}

func TestSearchCase2(t *testing.T) {
	actual := []int{-1, -100, 3, 99}
	target := 2
	expected := []int{3, 99, -1, -100}

	rotate(actual, target)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}

func TestSearchCase3(t *testing.T) {
	actual := []int{-1, -100, 3, 99}
	target := 3
	expected := []int{-100, 3, 99, -1}

	rotate(actual, target)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}

func TestSearchCase4(t *testing.T) {
	actualArr := []int{-1, -100, 3, 99}
	target := 0
	expectedArr := []int{-1, -100, 3, 99}

	rotate(actualArr, target)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d.", expected, actual)
		}
	}
}

func TestSearchCase5(t *testing.T) {
	actualArr := []int{-1}
	target := 213
	expectedArr := []int{-1}

	rotate(actualArr, target)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d.", expected, actual)
		}
	}
}

func TestSearchCase6(t *testing.T) {
	actualArr := []int{1, 2}
	target := 3
	expectedArr := []int{2, 1}

	rotate(actualArr, target)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("Expected: %d, Actual: %d.", expected, actual)
		}
	}
}

func TestSearchCase7(t *testing.T) {
	actualArr := []int{1, 2, 3, 4, 5, 6}
	target := 2
	expectedArr := []int{5, 6, 1, 2, 3, 4}

	rotate(actualArr, target)

	for i, actual := range actualArr {
		expected := expectedArr[i]
		if actual != expected {
			t.Errorf("\nExpected: %v,\nActual: %v.", expectedArr, actualArr)
			break
		}
	}
}

func TestSearchCase8(t *testing.T) {
	actual := []int{1, 2, 3, 4, 5, 6}
	target := 4
	expected := []int{3, 4, 5, 6, 1, 2}

	rotate(actual, target)

	for i, actualItem := range actual {
		expectedItem := expected[i]
		if actualItem != expectedItem {
			t.Errorf("\nExpected: %v,\nActual: %v.", expected, actual)
			break
		}
	}
}
