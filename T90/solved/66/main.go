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

	n := nextInt()
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}
	ans := 0.0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			cnt := 0
			for ii := l[i]; ii <= r[i]; ii++ {
				for jj := l[j]; jj <= r[j]; jj++ {
					if ii > jj {
						cnt++
					}
				}
			}
			ans += float64(cnt) / float64((r[i]-l[i]+1)*(r[j]-l[j]+1))
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
