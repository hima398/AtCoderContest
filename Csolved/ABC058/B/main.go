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
	o := nextString()
	e := nextString()
	ans := make([]string, len(o)+len(e))
	for i := 0; i < len(ans); i++ {
		if i%2 == 0 {
			ans[i] = string(o[i/2])
		} else {
			ans[i] = string(e[i/2])
		}
	}
	fmt.Println(strings.Join(ans, ""))
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
