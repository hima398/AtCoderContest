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
	ans := strings.Replace(s, "HAGIYA", "HAGIXILE", len(s))
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
