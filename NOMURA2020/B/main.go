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

	t := strings.Split(nextString(), "")
	ans := make([]string, len(t))
	for i := range t {
		if t[i] == "?" {
			ans[i] = "D"
		} else {
			ans[i] = t[i]
		}
	}
	fmt.Println(strings.Join(ans, ""))
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
