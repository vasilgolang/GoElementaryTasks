package task4

import (
	"testing"
)

//func TestNum2Digits(t *testing.T) {
//	var tests = []struct {
//		num    int
//		digits []uint8
//	}{
//		{num: 0, digits: []uint8{0, 0, 0, 0, 0, 0}},
//		{num: 1, digits: []uint8{0, 0, 0, 0, 0, 1}},
//		{num: 20, digits: []uint8{0, 0, 0, 0, 2, 0}},
//		{num: 100000, digits: []uint8{1, 0, 0, 0, 0, 0}},
//		{num: 555555, digits: []uint8{5, 5, 5, 5, 5, 5}},
//
//	}
//
//	for _, test := range tests {
//		if digits := num2Digits(test.num); reflect.DeepEqual(test.digits, digits) {
//			t.Errorf("num2Digits(%d) must return %v instead of %v", test.num, test.digits, digits)
//		}
//	}
//}

func TestFindMaxPalindrome(t *testing.T) {
	var tests = []struct {
		number     int
		palindrome int
		success    bool
	}{
		{
			number:     9123456789,
			palindrome: 0,
			success:    false,
		},
		{
			number:     11,
			palindrome: 11,
			success:    true,
		},
		{
			number:     121,
			palindrome: 121,
			success:    true,
		},
		{
			number:     1221,
			palindrome: 1221,
			success:    true,
		},
		{
			number:     81221,
			palindrome: 1221,
			success:    true,
		},
		{
			number:     12234444437,
			palindrome: 3444443,
			success:    true,
		},
	}

	for _, test := range tests {
		if palindrome, success := FindMaxPalindrome(test.number); success != test.success {
			t.Errorf("Success must be %v instead of %v for FindMaxPalindrome(%d)", test.success, success, test.number)
		} else {
			if palindrome != test.palindrome {
				t.Errorf("Palindrome for FindMaxPalindrome(%d) must be %d instead of %d", test.number, test.palindrome, palindrome)
			}
		}
	}
}

func BenchmarkFindMaxPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			FindMaxPalindrome(j)
		}
	}
}

func BenchmarkFindMaxPalindromeFromConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxPalindrome(123454321)
	}
}
