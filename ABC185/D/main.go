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
	if M == 0 {
		fmt.Println(1)
		return
	}
	A := make([]int, M+1)
	A[0] = 0
	for i := 1; i <= M; i++ {
		A[i] = nextInt()
	}
	sort.IntSlice(A).Sort()

	k := 9223372036854775807
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1]-1 == 0 {
			continue
		}
		k = Min(k, A[i]-A[i-1]-1)
	}
	// 白色のマスが無い
	if k == 9223372036854775807 {
		fmt.Println(0)
		return
	}
	var result int
	for i := 1; i < len(A); i++ {
		whites := A[i] - A[i-1] - 1
		if whites%k == 0 {
			result += whites / k
		} else {
			result += (whites / k) + 1
		}
	}
	if A[M] != N {
		whites := N - A[M]
		if whites%k == 0 {
			result += whites / k
		} else {
			result += (whites / k) + 1
		}
	}
	//	fmt.Println(k)
	fmt.Println(result)
}
