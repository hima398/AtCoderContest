package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const p = 200

var sc = bufio.NewScanner(os.Stdin)

func Format(s []int) (msg string) {
	msg = fmt.Sprintf("%d ", len(s))
	for i := range s {
		msg += fmt.Sprintf("%d ", s[i])
	}
	return
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	n2 := Min(n, 8) // 256
	max := (1 << n2) - 1

	//fmt.Printf("%b\n", max)
	memo := make([][]int, p)
	for i := 1; i <= max; i++ {
		sum := 0
		var s []int
		for j := 0; j < n2; j++ {
			if i&(1<<j) > 0 {
				s = append(s, j+1)
				sum += a[j]
				sum %= p
			}
		}
		if len(memo[sum]) > 0 {
			fmt.Println("Yes")
			fmt.Println(Format(memo[sum]))
			fmt.Println(Format(s))
			return
		} else {
			memo[sum] = s
		}
	}
	fmt.Println("No")
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
