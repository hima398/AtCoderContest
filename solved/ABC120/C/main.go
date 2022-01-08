package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := strings.Split(nextString(), "")
	n := len(s)
	var stk []string
	for i := 0; i < n; i++ {
		if len(stk) > 0 && stk[len(stk)-1] != s[i] {
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, s[i])
		}
	}
	//fmt.Println(stk)
	ans := n - len(stk)
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
