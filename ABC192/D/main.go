package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func F(x string, d int) int {
	n := len(x)
	sd := make([]int, n)
	sd[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if sd[i+1]*d/d != sd[i+1] {
			return -1
		}
		sd[i] = sd[i+1] * d
	}
	ret := 0
	for i := n - 1; i >= 0; i-- {
		xi := int(x[i] - '0')
		if xi*sd[i]/sd[i] != xi {
			return -1
		}
		if ret+xi*sd[i] < ret {
			return -1
		}
		ret += xi * sd[i]
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x := nextString()
	m := nextInt()

	if len(x) == 1 {
		xi := int(x[0] - '0')
		if xi <= m {
			// 値の種類なので1
			fmt.Println(1)
			return
		} else {
			// xi >= m
			fmt.Println(0)
			return
		}
	}

	d := 0
	for i := 0; i < len(x); i++ {
		xi := int(x[i] - '0')
		d = Max(d, xi)
	}

	if m < d {
		fmt.Println(0)
		return
	}
	l := d
	r := m + 1
	for r-l > 1 {
		mid := (r + l) / 2
		f := F(x, mid)
		if f < 0 {
			r = mid
		} else if f > m {
			//fmt.Println(l, r, f, m)
			r = mid
		} else if f <= m {
			// f < m
			//fmt.Println(l, r, f, m)
			l = mid
		}
	}
	fmt.Println(l - d)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
