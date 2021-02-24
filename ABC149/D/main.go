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

	n, k := nextInt(), nextInt()
	r, s, p := nextInt(), nextInt(), nextInt()
	t := nextString()
	mp := map[byte]int{'r': r, 's': s, 'p': p}
	mw := map[byte]byte{'r': 'p', 's': 'r', 'p': 's'}
	ml := map[byte]byte{'r': 's', 's': 'p', 'p': 'r'}
	ans := 0
	var tt []byte
	for i := 0; i < k; i++ {
		ans += mp[mw[t[i]]]
		tt = append(tt, mw[t[i]])
	}
	for i := k; i < n; i++ {
		if t[i] == t[i-k] && mw[t[i]] == tt[i-k] {
			tt = append(tt, ml[t[i]])
			continue
		}
		ans += mp[mw[t[i]]]
		tt = append(tt, mw[t[i]])
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
