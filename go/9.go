package main

import "fmt"

func isPalindrome12(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	a := x
	b := 0
	for a > 0 {
		tmp := a % 10
		a /= 10
		b = b*10 + tmp
	}

	return x == b
}

func main() {
	a := 232
	fmt.Println(isPalindrome12(a))
}
