package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveRegexp(s string) {
	r := regexp.MustCompile(`^(dream|dreamer|erase|eraser)+$`)
	if r.Match([]byte(s)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	SolveRegexp(s)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
