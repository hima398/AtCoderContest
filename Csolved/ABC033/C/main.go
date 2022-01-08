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

	s := nextString()
	ss := strings.Split(s, "+")
	nz := 0
	for _, str := range ss {
		if strings.Contains(str, "0") {
			nz++
		}
	}
	ans := len(ss) - nz
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
