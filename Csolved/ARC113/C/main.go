package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	m := make(map[byte]int)
	ans := 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == s[i-1] {
			// 置き換える文字数
			nr := len(s) - 1 - i
			if c, ok := m[s[i]]; ok {
				ans -= c
			}
			//fmt.Println(nr, m)
			ans += nr
			m = make(map[byte]int)
			m[s[i]] += nr + 2
			i--
		} else {
			m[s[i]]++
		}
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
