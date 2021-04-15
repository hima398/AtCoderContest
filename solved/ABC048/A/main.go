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
	ans := string(a[0]) + string(b[0]) + string(c[0])
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
