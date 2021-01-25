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

	m := map[string]string{"ABC": "ARC", "ARC": "ABC"}

	s := nextString()
	fmt.Println(m[s])
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
