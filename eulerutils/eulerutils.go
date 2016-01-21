package eulerutils

import (
	"math"
	"strconv"
)

func IsPrime(a int64) bool {
	if a <= 1 {
		return false
	}
	
	for b := Sqrt(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a % b == 0 {
			return false
		}
	}

	return true
}

func Sqrt(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func IsPalindrome(a int) bool {
	number := strconv.Itoa(a)
	var n int = len(number)
	var w int = 0

	for i := 0; i < (n / 2); i++ {
		w = (n - 1) - i
		if number[i] != number[w] {
			return false
		}
	}
	return true
}
