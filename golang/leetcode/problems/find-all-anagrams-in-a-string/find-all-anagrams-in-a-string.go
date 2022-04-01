package find_all_anagrams_in_a_string

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	sLength := len(s)
	pLength := len(p)
	var result []int

	if sLength < pLength {

		return result
	}

	smap := [26]int{}
	pmap := [26]int{}

	for i := 0; i < pLength; i++ {
		smap[s[i]-'a']++
		pmap[p[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if smap[i] == pmap[i] {
			count++
		}
	}

	//fmt.Printf("m1=%v\nm2=%v\n", pmap, smap)

	i := 0
	for ; i < sLength-pLength; i++ {
		right := s[i+pLength] - 'a'
		left := s[i] - 'a'

		if count == 26 {
			result = append(result, i)
		}

		smap[right]++
		if smap[right] == pmap[right] {
			count++
		} else if smap[right] == pmap[right]+1 {
			count--
		}
		smap[left]--
		if smap[left] == pmap[left] {
			count++
		} else if smap[left] == pmap[left]-1 {
			count--
		}
		//fmt.Printf("[%v] %v\nm1=%v\nm2=%v\n", i, count, pmap, smap)
	}

	if count == 26 {
		result = append(result, i)
	}

	return result
}
