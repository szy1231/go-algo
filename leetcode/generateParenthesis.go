package leetcode

//22. 括号生成

/*
	数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

    示例：

    输入：n = 3
    输出：[
    	   "((()))",
    	   "(()())",
    	   "(())()",
    	   "()(())",
    	   "()()()"
    	 ]

*/

//递归解法
func generateParenthesis(n int) []string {
	//如果数值非法，直接返回
	if n <= 0 {
		return nil
	}
	result := make([]string, 0)
	gen(0, 0, n, "", &result)
	return result
}

func gen(left, right, n int, s string, result *[]string) {
	//结束，满足条件，添加到切片中
	if left == right && left == n {
		result = append(result, s)
		return
	}
	if left < n {
		gen(left+1, right, n, s+"(", result)
	}
	//如果又括号数量 大于 左括号数量 是不合法的。剪枝直接剪掉
	if right < left {
		gen(left, right+1, n, s+")", result)
	}
}
