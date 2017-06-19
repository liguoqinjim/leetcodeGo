package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"abc", "abdc", "abd", "ab"}
	//strs := []string{"abc", "abdc", "abd"}
	//fmt.Println(longestCommonPrefix(strs))
	fmt.Println(longestCommonPrefix2(strs))
}

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

func longestCommonPrefix2(strs []string) string {
	fmt.Println(222)
	if strs == nil || len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	fmt.Println("pre1=", pre)
	for _, v := range strs {
		for !strings.HasPrefix(v, pre) {
			pre = pre[:len(pre)-1]
			fmt.Println("pre2=", pre)
		}
	}
	return pre
}
