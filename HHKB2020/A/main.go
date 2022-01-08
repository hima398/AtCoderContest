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
	t := nextString()
	if s == "Y" {
		fmt.Println(strings.ToUpper(t))
	} else { // s == "N"
		fmt.Println(t)
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
