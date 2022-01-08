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
	ss := strings.Split(s, "")
	t := nextString()
	n := len(s)
	for i := 0; i < n; i++ {
		if strings.Join(ss, "") == t {
			fmt.Println(i)
			return
		}
		//遅いけどlen(s) <= 1000程度なので間に合うはず
		ss = append([]string{string(ss[n-1])}, ss[:n-1]...)
	}
	fmt.Println(-1)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
