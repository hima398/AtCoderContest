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
	const max = 1000
	e := make([]int, max+1)
	for i := 2; i <= max; i++ {
		for j := i; j <= max; j += i {
			if e[j] == 0 {
				e[j] = i
			}
		}
	}
	n := nextInt()
	m := make(map[int]int)
	for i := n; i >= 2; i-- {
		idx := i
		for e[idx] != 0 {
			m[e[idx]]++
			idx /= e[idx]
		}
	}
	ans := 1
	for _, v := range m {
		ans *= v + 1
		ans %= p
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
