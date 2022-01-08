package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Format(t, m int) int {
	t1 := t % 10
	if m == 0 {
		if t1 >= 0 && t1 < 5 {
			return t - t1
		} else {
			return t - t1 + 5
		}
	} else if m == 1 {
		if t1 >= 1 && t1 < 5 {
			return t - t1 + 5
		} else if t1 > 5 && t1 <= 9 {
			return t - t1 + 10
		} else {
			return t
		}
	} else {
		return t
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s, e := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		se := nextString()
		fmt.Sscanf(se, "%d-%d", &s[i], &e[i])
		s[i], e[i] = Format(s[i], 0), Format(e[i], 1)
	}
	//fmt.Println(s)
	//fmt.Println(e)
	d := make([]int, 2401)
	for i := 0; i < n; i++ {
		d[s[i]]++
		d[e[i]]--
	}
	for i := 0; i < 2400; i++ {
		d[i+1] += d[i]
	}
	IsRainy := 0
	as := -1
	for i := 0; i <= 2400; i += 5 {
		if i%100 >= 60 {
			continue
		}
		if IsRainy == 0 && d[i] > 0 {
			as = i
			IsRainy ^= 1
		}
		if IsRainy == 1 && d[i] == 0 {
			fmt.Printf("%04d-%04d\n", as, i)
			as = -1
			IsRainy ^= 1
		}
	}
	if IsRainy == 1 {
		fmt.Printf("%04d-%04d\n", as, 2400)
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
