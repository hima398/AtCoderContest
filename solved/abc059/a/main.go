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

	s1, s2, s3 := nextString(), nextString(), nextString()
	ans := strings.ToUpper(string(s1[0])) + strings.ToUpper(string(s2[0])) + strings.ToUpper(string(s3[0]))
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
