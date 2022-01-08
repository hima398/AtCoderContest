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
	ss := strings.Replace(s, "25", "o", -1)
	//fmt.Println(ss)
	var p []int
	n := len(ss)
	cnt := 0
	for i := 0; i < n; i++ {
		if ss[i] == 'o' {
			cnt++
		} else {
			if cnt > 0 {
				p = append(p, cnt)
			}
			cnt = 0
		}
	}
	if cnt > 0 {
		p = append(p, cnt)
	}
	//fmt.Println(p)
	ans := 0
	F := func(x int) int {
		return x * (x + 1) / 2
	}
	for _, v := range p {
		ans += F(v)
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
