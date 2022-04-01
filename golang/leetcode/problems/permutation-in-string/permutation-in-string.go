package permutation_in_string

// https://leetcode.com/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	s1Length := len(s1)
	s2Length := len(s2)

	if s1Length > s2Length {
		return false
	}

	s1map := [26]int{}
	s2map := [26]int{}

	for i := 0; i < s1Length; i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1map[i] == s2map[i] {
			count++
		}
	}

	for i := 0; i < s2Length-s1Length; i++ {
		right := s2[i+s1Length] - 'a'
		left := s2[i] - 'a'

		if count == 26 {
			return true
		}

		s2map[right]++
		if s2map[right] == s1map[right] {
			count++
		} else if s2map[right] == s1map[right]+1 {
			count--
		}
		s2map[left]--
		if s2map[left] == s1map[left] {
			count++
		} else if s2map[left] == s1map[left]-1 {
			count--
		}
	}

	return count == 26
}
