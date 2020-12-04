package issubsequence

func isSubsequenceDualPointer(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0

	for i < len(s) && j < len(t) && i <= j {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	if i == len(s) {
		return true
	}
	return false
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	j := 0
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		var exist bool
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j < len(t) {
			exist = true
		}
		if i == 0 {
			dp[i] = exist
		} else {
			dp[i] = dp[i-1] && exist
		}
		j++
	}
	return dp[len(s)-1]
}
