package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(n int) int {
	m := make(map[string]int)
	for a := 1; a <= n; a++ {
		sa := strconv.Itoa(a)
		ab := string(sa[0]) + string(sa[len(sa)-1])
		if sa[len(sa)-1] != '0' {
			m[ab]++
		}
	}
	//fmt.Println(m)
	ret := 0
	for b := 1; b <= n; b++ {
		sb := strconv.Itoa(b)
		ba := string(sb[len(sb)-1]) + string(sb[0])
		ret += m[ba]
	}
	return ret
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	fmt.Println(Solve(n))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
