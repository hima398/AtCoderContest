package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = nextInt()
	}
	var ans []int
	l, r := 1, 1
	ok := true
	for ok && l < n {
		r = l
		for r <= n {
			if p[r] == l {
				break
			}
			r++
		}
		//fmt.Println(l, r)
		for i := r; i > l; i-- {
			ans = append(ans, i-1)
			p[i-1], p[i] = p[i], p[i-1]
			if i < r {
				if i != p[i] {
					ok = false
					break
				}
			}
		}
		//fmt.Println(p)
		l = r
	}
	if ok {
		for _, v := range ans {
			fmt.Fprintln(out, v)
			//fmt.Println(v)
		}
	} else {
		fmt.Fprintln(out, -1)
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
