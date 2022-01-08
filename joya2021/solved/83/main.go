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
	p := make([]int, n+1)
	ip := make([]int, n+1)
	for i := 1; i <= n; i++ {
		p[i] = nextInt()
		ip[p[i]] = i
	}
	//fmt.Println(p)
	//fmt.Println(ip)

	var ans []int
	used := make([]bool, n)
	used[0] = true
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	t := 1
	for t != n {
		for t != p[t] {
			i := ip[t] - 1
			p[i], p[i+1] = p[i+1], p[i]
			ip[p[i]], ip[p[i+1]] = ip[p[i+1]], ip[p[i]]
			if used[i] {
				fmt.Fprintln(out, -1)
				return
			} else {
				ans = append(ans, i)
				used[i] = true
			}
		}
		t++
	}
	//fmt.Println(used)
	allUsed := true
	for _, b := range used {
		allUsed = allUsed && b
	}
	if allUsed {
		for _, v := range ans {
			fmt.Fprintln(out, v)
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
