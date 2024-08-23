package longest_common_prefix

// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
