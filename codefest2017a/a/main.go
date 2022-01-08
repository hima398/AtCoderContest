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
	if strings.HasPrefix(s, "YAKI") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
