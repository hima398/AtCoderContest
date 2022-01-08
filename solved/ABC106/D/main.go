package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintCnt(cnt [][]int) {
	for _, c := range cnt {
		fmt.Println(c)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m, q := nextInt(), nextInt(), nextInt()
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		l, r := nextInt(), nextInt()
		l--
		r--
		cnt[l][r]++
	}
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			cnt[i][j] += cnt[i][j-1]
		}
	}
	//PrintCnt(cnt)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < q; i++ {
		p, q := nextInt(), nextInt()
		p--
		q--
		ans := 0

		for j := p; j <= q; j++ {
			ans += cnt[j][q]
		}
		fmt.Fprintln(out, ans)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
