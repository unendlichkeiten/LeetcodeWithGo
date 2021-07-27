package _009_palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	m, n := x, 0
	for m != 0 {
		n = 10*n + (m % 10)
		m /= 10
	}

	return x == n
}
