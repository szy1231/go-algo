package leetcode

//3. 无重复字符的最长子串

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return findMaxLen(s)
	}
}

func findMaxLen(s string) int {

	//maxLen:最大长度
	//curLen：当前长度
	//cur：当前坐标
	var maxLen, curLen, cur int

	//存储出现过的字符
	hashMap := make(map[int32]int)
	for k, v := range s {
		if n, ok := hashMap[v]; ok && n >= cur {

			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = k - (n + 1)
			cur = n + 1

		}
		hashMap[v] = k
		curLen++
	}
	//比较当前长度和最大长度
	if curLen > maxLen {
		return curLen
	}
	return maxLen

}
