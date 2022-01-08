package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := strings.Split(nextString(), "")
	n := nextInt()
	for i := 0; i < n; i++ {
		l, r := nextInt(), nextInt()
		l--
		r--
		for r-l > 0 {
			s[l], s[r] = s[r], s[l]
			r--
			l++
		}
	}
	ans := strings.Join(s, "")
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
