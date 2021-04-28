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

	const p = int(1e9) + 7
	_ = nextInt()
	s := nextString()

	m1 := map[byte]int{'a': 1, 't': 2, 'c': 3, 'o': 4, 'd': 5, 'e': 6, 'r': 7}
	m2 := make([]int, 8)
	idx := 0
	for i := range s {
		if m1[s[i]] == idx+1 {
			idx++
		}
		if m1[s[i]] == idx {
			m2[idx]++
		}
	}
	fmt.Println(m2)
	ans := 1
	for i := 1; i < 8; i++ {
		ans *= m2[i]
		ans %= p
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
