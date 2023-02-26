package main

import (
	"fmt"
	"strings"
)

func sqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func valid_palindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !check(s[i]) {
			i++
			continue
		}
		if !check(s[j]) {
			j--
			continue
		}

		if toLow(s[i]) != toLow(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func check(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func toLow(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return b
}

func valid_anagram(s, t string) bool {
	var countS = map[string]int{}
	var countT = map[string]int{}

	for _, v := range s {
		countS[string(v)]++
	}
	for _, v := range t {
		countT[string(v)]++
	}

	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}

func keyboard_row(words []string) []string {
	var result = []string{}
	var newMap = map[string]int{"q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
		"a": 2, "s": 2, "d": 2, "f": 2, "l": 2, "g": 2, "h": 2, "j": 2, "k": 2,
		"z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3}
	for _, v := range words {
		lower := strings.ToLower(string(v))
		var row = true
		for k := 1; k < len(lower); k++ {
			if newMap[string(lower[k])] != newMap[string(lower[k-1])] {
				row = false
			}
		}
		if row {
			result = append(result, string(v))
		}
	}
	return result
}

func main() {
	// task 1
	fmt.Println(sqrt(4)) // 2

	// task 2
	fmt.Println(valid_palindrome("A man, a plan, a canal: Panama")) // true

	// // task 3
	fmt.Println(valid_anagram("anagram", "nagaram")) // true

	// // task 4
	fmt.Println(keyboard_row([]string{"Hello", "Alaska", "Dad", "Peace"})) // ["Alaska","Dad"]
}
