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

	m := map[string]string{"1": "9", "9": "1"}
	n := nextString()
	for _, r := range n {
		fmt.Printf("%s", m[string(r)])
	}
	fmt.Println("")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
