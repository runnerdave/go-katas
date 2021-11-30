package main

func findPalindrome(s string) string {
	// 1- Create map with length as key and palindrome as value
	// 2- add first character
	// 3- select pivot as character in position 1 (n)
	// 4- compare characters in n-1 with n+1
	// 4.1- if equal add to map
	// 4.2- if not equal move pivot to n+1
	// 5- repeat 4 but compare to n=0
	// continue until pivot equals last value in the string select highest key as palindrome to return
	// only add the first palindrome to the map

	p := make(map[int]string)
	p[1] = string(s[0])
	for i := 1; i < len(s)-1; i++ {
		for j := 1; j < i+1 && len(s) > i+j; j++ {
			//two repeated letters
			if _, ok := p[2]; !ok && s[i] == s[i+1] {
				p[2] = s[i : i+2]
			}
			t := s[i-j : i+1+j]
			if _, ok := p[len(t)]; !ok && isPalindrome(t) {
				p[len(t)] = t
			}
		}
	}

	return p[returnLongestKey(p)]
}

func isPalindrome(s string) bool {
	h := len(s)
	for i := 0; i < h/2; i++ {
		if s[i] != s[h-1-i] {
			return false
		}
	}
	return true
}

func returnLongestKey(p map[int]string) int {
	k := 1
	for i := range p {
		if i > k {
			k = i
		}
	}
	return k
}
