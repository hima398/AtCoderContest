package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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

// return loser, winner
func VS(a []int) (int, int) {
	if len(a) == 2 {
		return Min(a[0], a[1]), Max(a[0], a[1])
	}
	_, w1 := VS(a[:len(a)/2])
	_, w2 := VS(a[len(a)/2:])
	return Min(w1, w2), Max(w1, w2)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	n = int(math.Pow(2.0, float64(n)))
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	l, _ := VS(a)

	for i, v := range a {
		if v == l {
			fmt.Println(i + 1)
			return
		}
	}
}
