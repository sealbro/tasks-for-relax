package number_of_islands

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	grid := [][]byte{
		{'0', '0', '1', '0', '0', '0', '0', '1', '0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0', '0', '0'},
		{'0', '1', '1', '0', '1', '0', '0', '0', '0', '0', '0', '0', '0'},
		{'0', '1', '0', '0', '1', '1', '0', '0', '1', '0', '1', '0', '0'},
		{'0', '1', '0', '0', '1', '1', '0', '0', '1', '1', '1', '0', '0'},
		{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '1', '0', '0'},
		{'0', '0', '0', '0', '0', '0', '0', '1', '1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0', '0', '0', '1', '1', '0', '0', '0', '0'},
	}

	expected := 6

	actual := numIslands(grid)

	assert.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	grid := [][]byte{
		{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'},
	}

	expected := 0

	actual := numIslands(grid)

	assert.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '1'},
		{'0', '0', '0', '1', '1'},
	}

	expected := 2

	actual := numIslands(grid)

	assert.Equal(t, expected, actual)
}
func TestCase4(t *testing.T) {
	grid := [][]byte{

		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	expected := 1

	actual := numIslands(grid)

	assert.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	expected := 3

	actual := numIslands(grid)

	assert.Equal(t, expected, actual)
}
