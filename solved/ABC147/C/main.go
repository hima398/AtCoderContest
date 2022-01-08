package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	ln := nextInt()
	a := make([]int, ln)
	m := make([]map[int]int, ln)
	for i := 0; i < ln; i++ {
		m[i] = make(map[int]int)
	}
	for i := 0; i < ln; i++ {
		a[i] = nextInt()
		for j := 0; j < a[i]; j++ {
			x, y := nextInt(), nextInt()
			x--
			m[i][x] = y
		}
	}

	n := (1 << ln) - 1
	ans := 0
	for mask := n; mask >= 0; mask-- {
		/*
			for j := 0; j < ln; j++ {
				d |= i & (1 << j)
			}
		*/
		// 証言に基づく正直者リスト
		ok := true
		for j := 0; j < ln; j++ {
			if (mask>>j)&1 == 1 {
				for k := 0; k < ln; k++ {
					// 証言が存在しない場合
					if _, ok := m[j][k]; !ok {
						continue
					}
					// 証言に矛盾がある場合
					if m[j][k] != (mask>>k)&1 {
						ok = false
					}
				}
			}
		}
		if ok {
			//fmt.Printf("%b\n", mask)
			ans = Max(ans, bits.OnesCount(uint(mask)))
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
