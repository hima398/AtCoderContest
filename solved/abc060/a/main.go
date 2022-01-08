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

	a, b, c := nextString(), nextString(), nextString()
	na, nb := len(a), len(b)
	if a[na-1] == b[0] && b[nb-1] == c[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
