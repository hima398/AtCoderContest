package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_, m := nextInt(), nextInt()
	a := make([]int, m)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	k := nextInt()
	c := make([]int, k)
	d := make([]int, k)
	ball := make(map[int]int)
	for i := 0; i < k; i++ {
		c[i], d[i] = nextInt(), nextInt()
		ball[c[i]]++
		ball[d[i]]++
	}
	//fmt.Println(ball)
	ans := 0
	pattern := 1
	for pattern < int(math.Pow(2.0, float64(k)))-1 {
		dishes := make(map[int]int)
		for i := 0; i < k; i++ {
			if pattern>>i&1 == 1 {
				dishes[c[i]]++
			} else {
				dishes[d[i]]++
			}
		}

		sum := 0
		for j := 0; j < m; j++ {
			if dishes[a[j]] > 0 && dishes[b[j]] > 0 {
				sum++
			}
		}
		ans = Max(ans, sum)
		pattern++
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
