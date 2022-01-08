package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestry(n int, t []int) int {
	var ans int
	if n == 1 {
		ans = t[0]
	} else if n == 2 {
		ans = Max(t[0], t[1])
	} else if n == 3 {
		ans = Max(t[0]+t[1], t[2])
		ans = Min(ans, Max(t[0]+t[2], t[1]))
		ans = Min(ans, Max(t[1]+t[2], t[0]))
	} else {
		ans = Max(t[0]+t[1], t[2]+t[3])
		ans = Min(ans, Max(t[0]+t[2], t[1]+t[3]))
		ans = Min(ans, Max(t[0]+t[3], t[1]+t[2]))
		ans = Min(ans, Max(t[0]+t[1]+t[2], t[3]))
		ans = Min(ans, Max(t[0]+t[1]+t[3], t[2]))
		ans = Min(ans, Max(t[0]+t[2]+t[3], t[1]))
		ans = Min(ans, Max(t[1]+t[2]+t[3], t[0]))
	}
	return ans
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	t := nextIntSlice(n)

	fmt.Println(SolveHonestry(n, t))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
