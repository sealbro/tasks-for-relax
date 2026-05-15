package find_all_possible_recipes_from_given_supplies

import (
	"golang/assert"
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	recipes := []string{"bread"}
	ingredients := [][]string{{"yeast", "flour"}}
	supplies := []string{"yeast", "flour", "corn"}

	expected := []string{"bread"}

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}

func TestCase2(t *testing.T) {
	recipes := []string{"bread", "sandwich"}
	ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}}
	supplies := []string{"yeast", "flour", "meat"}

	expected := []string{"bread", "sandwich"}

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}

func TestCase3(t *testing.T) {
	recipes := []string{"bread", "sandwich", "burger"}
	ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}}
	supplies := []string{"yeast", "flour", "meat"}

	expected := []string{"bread", "burger", "sandwich"}

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}

func TestCase4(t *testing.T) {
	recipes := []string{"sandwich", "burger"}
	ingredients := [][]string{{"bread", "meat"}, {"sandwich", "meat", "bread"}}
	supplies := []string{"yeast", "flour", "meat"}

	var expected []string

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}

func TestCase5(t *testing.T) {
	recipes := []string{"bread", "sandwich", "burger", "kebab"}
	ingredients := [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}, {"meat", "flour", "spicy"}}
	supplies := []string{"yeast", "flour", "meat"}

	expected := []string{"bread", "burger", "sandwich"}

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}

func TestCase6(t *testing.T) {
	recipes := []string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}
	ingredients := [][]string{
		{"d"},                   // ju +
		{"hveml", "f", "cpivl"}, // fzjnm +
		{"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"},           // x
		{"cpivl", "hveml", "zpmcz", "ju", "h"},                // e
		{"h", "fzjnm", "e", "q", "x"},                         // zpmcz
		{"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"}, // h
		{"f", "hveml", "cpivl"}}                               // q +
	supplies := []string{"f", "hveml", "cpivl", "d"}

	expected := []string{"fzjnm", "ju", "q"}

	actual := findAllRecipes(recipes, ingredients, supplies)

	slices.Sort(actual)
	assert.EqualMany(t, expected, actual)
}
