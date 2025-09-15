package main

/* 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
   回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
   • 例如，121 是回文，而 123 不是。*/

func getPalindromicNumber(i int) bool {

	//当 i < 0 时，i 不是回文数。
	if i < 0 || (i%10 == 0 && i != 0) {
		return false
	}
	reverteNumber := 0

	for i > reverteNumber {
		reverteNumber = reverteNumber*10 + i%10
		i /= 10
	}

	if i == reverteNumber || i == reverteNumber/10 {
		return true
	} else {
		return false
	}

}
