package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N, M := nextInt(), nextInt()
	A := make([]int, M+1)
	A[0] = 0
	for i := 1; i <= M; i++ {
		A[i] = nextInt()
	}
	sort.IntSlice(A).Sort()
	k := 9223372036854775807
	for i := 1; i <= M; i++ {
		if A[i]-A[i-1]-1 == 0 {
			continue
		}
		k = Min(k, A[i]-A[i-1]-1)
	}
	if k == 0 {
		fmt.Println(0)
		return
	}
	var result int64
	for i := 1; i <= M; i++ {
		result += int64(float64(A[i]-A[i-1]-1) + 0.5/float64(k))
	}
	result += int64(float64(N-A[M]-1) + 0.5/float64(k))
	fmt.Println(k)
	fmt.Println(result)
}
