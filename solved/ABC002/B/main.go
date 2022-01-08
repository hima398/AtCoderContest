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

	w := nextString()
	w = strings.ReplaceAll(w, "a", "")
	w = strings.ReplaceAll(w, "i", "")
	w = strings.ReplaceAll(w, "u", "")
	w = strings.ReplaceAll(w, "e", "")
	w = strings.ReplaceAll(w, "o", "")
	fmt.Println(w)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
