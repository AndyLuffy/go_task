package main

/*
字符串
• 有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func getValidParentheses(s string) bool {

	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < n; i += 2 {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0

}
