package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestry(a, b int) int {

	ans := 0
	for i := a; i <= b; i++ {
		ans ^= i
	}
	return ans
}

func F(n int) int {
	if n%2 == 0 {
		if (n/2)%2 == 0 {
			return n
		} else {
			return 1 ^ n
		}
	} else {
		if Ceil(n, 2)%2 == 0 {
			return 0
		} else {
			return 1
		}
	}
}

func Solve(a, b int) int {
	xb := F(b)
	if a == 0 {
		return xb
	}
	xa := F(a - 1)
	//fmt.Printf("%b, %b\n", xa, xb)
	return xb ^ xa
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	//fmt.Println(SolveHonestry(a, b))
	fmt.Println(Solve(a, b))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
