package leetcode

import "strings"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/

//1.利用栈判断
type Stack []string

func (s *Stack) Pop() string {
	ns := *s
	v := ns[len(ns)-1]
	*s = ns[:len(ns)-1]
	return v
}
func isValid(s string) bool {

	if s == "" {
		return true
	}

	var stack Stack
	hashMap := map[string]string{
		"]": "[",
		")": "(",
		"}": "{",
	}

	for _, v := range s {
		if _, ok := hashMap[string(v)]; !ok {
			stack = append(stack, string(v))
		} else if len(stack) == 0 || stack.Pop() != hashMap[string(v)] {
			return false
		}
	}
	return len(stack) == 0
}

//2.替换所有括号
func isValid02(s string) bool {

	if s == "" {
		return true
	}

	for {
		old := s
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "()", "")

		if s == old {
			return false
		}
		if s == "" {
			return true
		}
	}

}
