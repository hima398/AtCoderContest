package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Section struct {
	l, r int
}

func F1(n, m int) (ans []Section) {
	for i := 0; i < n; i++ {
		ans = append(ans, Section{2*i + 1, 2 * (i + 1)})
	}
	return ans
}

func F2(n, m int) (ans []Section) {
	l := 2
	for i := 0; i <= m; i++ {
		ans = append(ans, Section{l, l + 1})
		l += 2
	}
	ans = append(ans, Section{1, l})
	l++
	for i := 0; i < n-m-2; i++ {
		ans = append(ans, Section{l, l + 1})
		l += 2
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()

	var ans []Section
	if m == 0 {
		ans = F1(n, m)
	} else if m > 0 && m <= n-2 {
		ans = F2(n, m)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if len(ans) > 0 {
		for _, v := range ans {
			fmt.Fprintln(out, v.l, v.r)
		}
	} else {
		//if m < 0 || m == n || (m >= n-1 && n >= 2) {
		fmt.Fprintln(out, -1)
	}

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
