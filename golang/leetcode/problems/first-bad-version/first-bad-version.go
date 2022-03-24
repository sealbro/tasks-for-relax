package first_bad_version

// https://leetcode.com/problems/first-bad-version/

func firstBadVersion(n int) int {
	findLeft := 0
	findRight := n

	var pivot int
	for findLeft+1 != findRight {
		pivot = findLeft + (findRight-findLeft)/2

		if isBadVersion(pivot) {
			findRight = pivot
		} else {
			findLeft = pivot
		}
	}

	return findLeft + 1
}

func isBadVersion(version int) bool {
	return version == 4
}
