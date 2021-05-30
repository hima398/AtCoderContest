package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Dog struct {
	a int
	c string
}

func SearchMinDiff(a, b []Dog) int {
	const INF = int(1e15) + 1
	if len(a) == 0 || len(b) == 0 {
		return INF
	}
	var ts []Dog
	ts = append(ts, a...)
	ts = append(ts, b...)
	sort.Slice(ts, func(i, j int) bool { return ts[i].a < ts[j].a })
	ret := INF
	for i := 0; i < len(ts)-1; i++ {
		if ts[i+1].c != ts[i].c {
			ret = Min(ret, ts[i+1].a-ts[i].a)
		}
	}
	return ret
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	nh := nextInt()
	nd := 2 * nh
	//dogs := make([]Dog, nd)
	//var dogs [3][]int
	var rd, gd, bd []Dog
	for i := 0; i < nd; i++ {
		a, c := nextInt(), nextString()
		//dogs[i].a, dogs[i].c = nextInt(), nextString()
		switch c {
		case "R":
			rd = append(rd, Dog{a, c})
		case "G":
			gd = append(gd, Dog{a, c})
		default:
			// c == "B"
			bd = append(bd, Dog{a, c})
		}
	}
	//fmt.Println(rd, gd, bd)
	if len(rd)%2 == 0 && len(gd)%2 == 0 && len(bd)%2 == 0 {
		fmt.Println(0)
		return
	}
	var ans int
	if len(rd)%2 == 0 {
		// R == even, G, B == odd
		ans = Min(SearchMinDiff(gd, bd), SearchMinDiff(rd, gd)+SearchMinDiff(rd, bd))
	} else if len(gd)%2 == 0 {
		// G == even, R, B == odd
		ans = Min(SearchMinDiff(rd, bd), SearchMinDiff(rd, gd)+SearchMinDiff(gd, bd))

	} else {
		// B == even, G, B == odd
		ans = Min(SearchMinDiff(rd, gd), SearchMinDiff(rd, bd)+SearchMinDiff(gd, bd))
	}
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
