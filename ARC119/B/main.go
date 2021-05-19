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

	// 01111111
	// 11111110
	n := nextInt()
	s := nextString()
	t := nextString()

	var ss []int
	var tt []int
	for i := 0; i < n; i++ {
		if int(s[i]-'0') == 0 {
			ss = append(ss, i)
		}

		if int(t[i]-'0') == 0 {
			tt = append(tt, i)
		}
	}
	if len(ss) != len(tt) {
		fmt.Println(-1)
		return
	}
	ans := 0
	for i := range ss {
		if ss[i] != tt[i] {
			ans++
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
