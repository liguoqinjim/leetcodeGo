package lab014

import "strings"

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
