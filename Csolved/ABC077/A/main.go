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

	c1 := nextString()
	c2 := nextString()
	if c1[0] == c2[2] && c1[1] == c2[1] && c1[2] == c2[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
