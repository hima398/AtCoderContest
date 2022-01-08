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

	n, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
	l := nextIntSlice(n)
	mask := 1 << (2 * n)
	ans := int(1e5)
	for k := 0; k < mask; k++ {
		mp := 0
		var la, lb, lc int
		for i := 0; i < n; i++ {
			switch k >> (2 * i) & 3 {
			case 1:
				if la != 0 {
					mp += 10
				}
				la += l[i]
			case 2:
				if lb != 0 {
					mp += 10
				}
				lb += l[i]
			case 3:
				if lc != 0 {
					mp += 10
				}
				lc += l[i]
			default:
			}
		}
		mp += Abs(a-la) + Abs(b-lb) + Abs(c-lc)
		//1110101001
		/*
			if k == 0x3a9 {
				fmt.Printf("%b ", k)
				fmt.Println(la, lb, lc, mp)
			}
		*/
		//a: 400, 600 b: 300, 500, c:
		if la != 0 && lb != 0 && lc != 0 {
			ans = Min(ans, mp)
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
