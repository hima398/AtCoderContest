package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	sl := strings.ToLower(s)
	reg := regexp.MustCompile(`i[a-z]*c[a-z]*t`)
	if reg.Match([]byte(sl)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
