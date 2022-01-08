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
	n2 := n / 2
	isOdd := n%2 == 1
	nd := 0
	if n == 1 {
		fmt.Println(0)
		return
	}

	for i := 0; i < n2; i++ {
		if s[i] != s[n-i-1] {
			nd++
		}
	}

	ans := 25 * n
	if nd > 1 {
		fmt.Println(ans)
	} else if nd == 1 {
		ans -= 2
		fmt.Println(ans)
	} else {
		// nd == 0
		if isOdd {
			ans -= 25
			fmt.Println(ans)
		} else {
			fmt.Println(ans)
		}
	}

}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
