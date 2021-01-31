package main

import (
	"bufio"
	"errors"
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

func WithMap(N int) {
	M := make(map[int]bool)
	for i := 70; i < 80; i++ {
		M[i] = true
	}
	for i := 700; i < 800; i++ {
		M[i] = true
	}
	for i := 7000; i < 8000; i++ {
		M[i] = true
	}
	for i := 70000; i < 80000; i++ {
		M[i] = true
	}
	for i := 070; i <= 077; i++ {
		M[i] = true
	}
	for i := 0700; i <= 0777; i++ {
		M[i] = true
	}
	for i := 07000; i <= 07777; i++ {
		M[i] = true
	}
	for i := 070000; i < 077777; i++ {
		M[i] = true
	}
	for i := 0700000; i < 0777777; i++ {
		M[i] = true
	}

	ans := 0
	for i := 1; i <= N; i++ {
		if i%8 == 7 {
			continue
		}
		if i%10 == 7 {
			continue
		}
		if M[i] {
			continue
		}
		ans++
	}
	fmt.Println(M)
	fmt.Println(ans)

}

func ToSlice(N int) ([]int, []int) {
	var D []int
	T1 := N
	for T1 > 0 {
		D = append(D, T1%10)
		T1 /= 10
	}
	//fmt.Println(D)
	var E []int
	T2 := N
	for T2 > 0 {
		E = append(E, T2%8)
		T2 /= 8
	}
	//fmt.Println(E)
	return D, E
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	N := nextInt()
	ans := 0
	for i := 1; i <= N; i++ {
		D, E := ToSlice(i)
		Dhas7 := false
		for _, v := range D {
			if v == 7 {
				Dhas7 = true
				break
			}
		}
		Ehas7 := false
		for _, v := range E {
			if v == 7 {
				Ehas7 = true
				break
			}
		}
		if !Dhas7 && !Ehas7 {
			ans++
		}
	}
	fmt.Println(ans)
}
