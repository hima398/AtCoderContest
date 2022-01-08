package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextInt()
	if t == 1 {
		n := len(s)
		ans := 0
		x, y := 0, 0
		for i := 0; i < n; i++ {
			switch s[i] {
			case 'L':
				x--
			case 'R':
				x++
			case 'U':
				y++
			case 'D':
				y--
			case '?':
				ans++
			}
		}
		ans += Abs(x) + Abs(y)
		fmt.Println(ans)
	} else {
		// t == 2
		n := len(s)
		x, y := 0, 0
		nq := 0
		for i := 0; i < n; i++ {
			switch s[i] {
			case 'L':
				x--
			case 'R':
				x++
			case 'U':
				y++
			case 'D':
				y--
			case '?':
				nq++
			}
		}
		ans := Abs(x) + Abs(y)
		if ans > nq {
			ans -= nq
		} else {
			// ans =< nq
			nq -= ans
			if nq%2 == 0 {
				ans = 0
			} else {
				ans = 1
			}
		}
		fmt.Println(ans)
	}
}

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
