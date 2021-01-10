package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m := map[string]int{
		"SSS": 0, "SSR": 1, "SRS": 1, "SRR": 2,
		"RSS": 1, "RSR": 1, "RRS": 2, "RRR": 3}
	s := nextString()
	fmt.Println(m[s])
}
