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

	n, s := nextInt(), nextString()

	ans := 0
	for i := 0; i < n; i++ {
		m := make(map[byte]int)
		for j := i; j < n; j++ {
			m[s[j]]++
			//fmt.Println(m)
			if m['A'] == m['T'] && m['C'] == m['G'] {
				ans++
			}
		}
	}
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
