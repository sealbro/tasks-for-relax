package verifying_an_alien_dictionary

import (
	"golang/assert"
	"testing"
)

func TestCase1(t *testing.T) {
	words := []string{"hello", "leetcode"}
	target := "hlabcdefgijkmnopqrstuvwxyz"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase2(t *testing.T) {
	words := []string{"word", "world", "row"}
	target := "worldabcefghijkmnpqstuvxyz"

	actual := isAlienSorted(words, target)

	assert.False(t, actual)
}

func TestCase3(t *testing.T) {
	words := []string{"apple", "app"}
	target := "abcdefghijklmnopqrstuvwxyz"

	actual := isAlienSorted(words, target)

	assert.False(t, actual)
}

func TestCase4(t *testing.T) {
	words := []string{"he", "leetcode"}
	target := "hlabcdefgijkmnopqrstuvwxyz"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase5(t *testing.T) {
	words := []string{"he", "le"}
	target := "hlabcdefgijkmnopqrstuvwxyz"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase6(t *testing.T) {
	words := []string{"hello", "lee"}
	target := "hlabcdefgijkmnopqrstuvwxyz"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase7(t *testing.T) {
	words := []string{"kuvp", "q"}
	target := "ngxlkthsjuoqcpavbfdermiywz"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase8(t *testing.T) {
	words := []string{"ubg", "kwh"}
	target := "qcipyamwvdjtesbghlorufnkzx"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase9(t *testing.T) {
	words := []string{"qwerty", "qwertya", "qwertyav"}
	target := "qcipyamwvdjtesbghlorufnkzx"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}

func TestCase11(t *testing.T) {
	words := []string{"vwtvtkwcxsdxxcjiq", "slssdoehjkgmtilyn", "stkmrqmbozeypxrrg", "duikzoxxspomrsfwr", "rlkmgtqbhuxcjsupw", "mgkxwmiykftnvwxaa", "nsnrozwctmqzouwva", "emhbrphsgbhfwjbwr", "ojxhvtgomwidsvyxh", "ehrimaakwdjikbhte", "bsjsbhlqsimaohsqf", "nytkwpivkqctcwxto", "iupxpavankliwavgy", "tvqyftrkmscqwmvrg", "gyosqifvrwjelwytx", "evmiiwlzuftupszih", "ycfmoppkfcbmaghhs", "ocanttjfpakexymok", "hymlhxrhglrlapvcu", "plnbguunqekelbzpy", "vujinbnkmunpjnkvt", "saxiujwfhlsavbrln", "xpxmqqngrgnnrngel", "etxhrdwozeeopjirz", "pqwngrjbogtliyxdi", "gzpugidmfvohixmsb", "wyqulfrceppctskyf", "mukykfxkgahwpjfxy", "ugocttswohsneckwz", "pinlpohhzjmmekclr", "krnfnlgobbiorzect", "jsmcartuvqixdfbkd", "zhodrwubunmxsyeoe", "mvwofemierzwnxxgj", "ohlgzickqheosvbao", "ymyuinbrapqakxcea", "rtgkezktviqzadlqs", "qyjyfzifznumvuudk", "kvugcorqyixsbjxhj", "btfoxxcadmuwwvcqd"}
	target := "zavfwhlsjomuebitdkcqnyrxpg"

	actual := isAlienSorted(words, target)

	assert.False(t, actual)
}

func TestCase12(t *testing.T) {
	words := []string{"qwerty", "qwerty", "qwerty"}
	target := "qcipyamwvdjtesbghlorufnkzx"

	actual := isAlienSorted(words, target)

	assert.True(t, actual)
}
