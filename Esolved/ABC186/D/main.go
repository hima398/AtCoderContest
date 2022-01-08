package main

import (
	"bufio"
	"errors"
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

func DivideSlice(A []int, K int) ([]int, []int, error) {

	if len(A) < K {
		return nil, nil, errors.New("")
	}
	return A[:K+1], A[K:], nil
}

type Pos struct {
	X int
	Y int
}

func Stupid(N int, A []int) int {
	sum := 0
	for i := 1; i < N; i++ {
		for j := i + 1; j <= N; j++ {
			sum += Abs(A[i] - A[j])
		}
	}
	return sum
}

func Solve(N int, A, S []int) int {
	sum := 0

	for i := 0; i+1 < N; i++ {
		sum += (S[N] - S[i+1]) - (N-i-1)*A[i]
	}
	return sum
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	A := make([]int, N)
	S := make([]int, N+1)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	sort.Ints(A)
	S[0] = 0
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}

	//fmt.Println(Stupid(N, A))
	fmt.Println(Solve(N, A, S))
}
