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
	t := nextString()
	n := len(s) // len(s) == len(t)
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)

	for i := 0; i < n; i++ {
		_, ok := ms[s[i]]
		if ok {
			// 変換表に矛盾がある
			if ms[s[i]] != t[i] {
				fmt.Println("No")
				return
			}
		}
		ms[s[i]] = t[i]
	}
	for i := 0; i < n; i++ {
		_, ok := mt[t[i]]
		if ok {
			// 変換表に矛盾がある
			if mt[t[i]] != s[i] {
				fmt.Println("No")
				return
			}
		}
		mt[t[i]] = s[i]
	}
	fmt.Println("Yes")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
