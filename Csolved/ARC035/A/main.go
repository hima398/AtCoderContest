package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	n := len(s)
	isPalindrome := true
	for i := 0; i < n/2; i++ {
		isPalindrome = isPalindrome && (s[i] == s[n-i-1] || s[i] == '*' || s[n-i-1] == '*')
	}
	if isPalindrome {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
