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

	s := []string{"1", "2", "3", "4", "5", "6"}

	p := make([]string, 30)
	// 36パターン並びを作っておく
	for i := 0; i < 30; i++ {
		p[i] = strings.Join(s, "")
		//fmt.Println(i, p2[i])
		ii, jj := i%5, i%5+1
		s[ii], s[jj] = s[jj], s[ii]
	}

	n := nextInt()

	ans := p[n%30]
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
