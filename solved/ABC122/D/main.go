package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	const p = int(1e9) + 7

	n := nextInt()

	memo := make([]map[string]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make(map[string]int, n)
	}

	IsOk := func(rear string) bool {
		for i := 0; i < 4; i++ {
			s := strings.Split(rear, "")
			if i >= 1 {
				s[i-1], s[i] = s[i], s[i-1]
			}
			if strings.Contains(strings.Join(s, ""), "AGC") {
				return false
			}
		}
		return true
	}

	var F func(i int, rear string) int
	F = func(i int, rear string) int {
		if _, ok := memo[i][rear]; ok {
			return memo[i][rear]
		}
		if i == n {
			return 1
		}
		ret := 0
		for _, c := range "ACGT" {
			if IsOk(rear + string(c)) {
				ret += F(i+1, rear[1:]+string(c))
				ret %= p
			}
		}
		memo[i][rear] = ret
		return ret
	}

	ans := F(0, "###")
	/*
		for _, v := range memo {
			fmt.Println(v)
		}
	*/
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
