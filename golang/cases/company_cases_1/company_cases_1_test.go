package company_cases_1

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	x, y, z := Case1()

	expected := []int{0, 1, 2, 4}

	assert.EqualMany(t, expected[:3], x)
	assert.EqualMany(t, expected, y)
	assert.EqualMany(t, expected, z)
}

func TestCase2(t *testing.T) {
	seconds := Case2()

	assert.Equal(t, 2, seconds)
}

func TestCase3(t *testing.T) {
	withoutZeros := Case3()

	assert.EqualMany(t, []int{2, 1, 1}, withoutZeros)
}

func TestCase4(t *testing.T) {
	count := Case4()

	assert.Equal(t, 100, count)
}

func TestCase5(t *testing.T) {
	list := &Item[int]{
		Value: 0,
		Next: &Item[int]{
			Value: 1,
			Next: &Item[int]{
				Value: 2,
				Next: &Item[int]{
					Value: 3,
				},
			},
		},
	}

	list = Reverse(list)

	expected := []int{3, 2, 1, 0}

	var actual []int
	cur := list
	for cur != nil {
		actual = append(actual, cur.Value)
		cur = cur.Next
	}

	assert.EqualMany(t, expected, actual)
}
