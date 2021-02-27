package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeScore(s string) int {
	hand := make([]int, 10)
	for i := 0; i < len(s); i++ {
		hand[int(s[i]-'0')]++
	}
	ret := 0
	for i := 1; i < len(hand); i++ {
		ret += i * int(math.Pow10(hand[i]))
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	k := nextInt()
	s := nextString()
	t := nextString()
	cards := make([]int, 10)
	for i := 1; i < 10; i++ {
		cards[i] = k
	}
	tk := make([]int, 10)
	ao := make([]int, 10)
	for i := 0; i < 4; i++ {
		tk[int(s[i]-'0')]++
		ao[int(t[i]-'0')]++
		cards[int(s[i]-'0')]--
		cards[int(t[i]-'0')]--
	}
	//fmt.Println(tk)
	//fmt.Println(ao)

	win := 0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			ss := s[:4] + strconv.Itoa(i)
			tt := t[:4] + strconv.Itoa(j)
			if ComputeScore(ss) <= ComputeScore(tt) {
				continue
			}
			if i == j {
				win += cards[i] * (cards[i] - 1)
			} else {
				win += cards[i] * cards[j]
			}
		}
	}
	rem := 9*k - 8
	ans := float64(win) / float64(rem) / float64(rem-1)
	fmt.Println(ans)
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
