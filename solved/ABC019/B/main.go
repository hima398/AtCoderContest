package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	c := s[0]
	cnt := 1
	ans := ""
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			ans += string(c) + strconv.Itoa(cnt)
			c = s[i+1]
			cnt = 1
		} else {
			cnt++
		}
	}
	ans += string(c) + strconv.Itoa(cnt)
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
