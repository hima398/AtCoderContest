package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n, m int) int {
	k := make([]int, m)
	s := make([][]int, m)
	for i := 0; i < m; i++ {
		k[i] = nextInt()
		//fmt.Println(k[i])
		ss := make([]int, k[i])
		for j := 0; j < k[i]; j++ {
			ss[j] = nextInt()
		}
		s[i] = ss
		//fmt.Println(ss)
	}
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = nextInt()
	}
	max := 1 << n
	ans := 0
	for pat := 0; pat < max; pat++ {
		isAllOn := true
		for i := 0; i < m; i++ {
			on := 0
			for j := 0; j < k[i]; j++ {
				if pat&(1<<(s[i][j]-1)) != 0 {
					on++
				}
			}
			isAllOn = isAllOn && (on%2 == p[i])
		}
		if isAllOn {
			ans++
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	fmt.Println(SolveHonestly(n, m))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
