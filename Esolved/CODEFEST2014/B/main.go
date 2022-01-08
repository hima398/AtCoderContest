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
	ans := int(s[0] - '0')
	n := len(s)
	for i := 1; i < n; i++ {
		v := int(s[i] - '0')
		if i%2 == 1 {
			ans -= v
		} else {
			ans += v
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
