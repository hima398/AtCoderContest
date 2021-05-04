package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func IsPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	n := len(s)
	if IsPalindrome(s) && IsPalindrome(s[0:n/2]) && IsPalindrome(s[n/2+1:]) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
