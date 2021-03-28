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

	s := nextString()
	sd := s[1:] + string(s[0])
	fmt.Println(sd)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
