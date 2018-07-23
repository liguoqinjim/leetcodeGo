package lab0014

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	n := 0
	end := false
	for !end {
		for i := 0; i < len(strs)-1; i++ {
			if n == len(strs[i]) || n == len(strs[i+1]) {
				end = true
				break
			}
			if strs[i][n] != strs[i+1][n] {
				end = true
				break
			}
		}
		if !end {
			n++
		}
	}

	return strs[0][:n]
}

//sample
func longestCommonPrefix2(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for _, v := range strs {
		for !strings.HasPrefix(v, pre) {
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
