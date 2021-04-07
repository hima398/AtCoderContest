package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	n := len(s)
	ans := string(s[0]) + strconv.Itoa(n-2) + string(s[n-1])
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
