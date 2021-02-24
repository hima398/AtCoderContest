package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

type Pos struct {
	X int
	Y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N, M, T := nextInt(), nextInt(), nextInt()
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		A[i] = nextInt()
		B[i] = nextInt()
	}
	V := N
	Now := 0
	for i := 0; i < M; i++ {
		V -= (A[i] - Now)
		//		fmt.Println(V)
		if V <= 0 {
			fmt.Println("No")
			return
		}
		V = Min(V+(B[i]-A[i]), N)
		Now = B[i]
	}
	V -= (T - B[M-1])
	if V <= 0 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	return
}
