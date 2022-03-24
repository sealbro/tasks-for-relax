package first_bad_version

import "testing"

func TestSearchCase1(t *testing.T) {
	n := 5
	expected := 4 // bad

	actual := firstBadVersion(n)

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d.", expected, actual)
	}
}
