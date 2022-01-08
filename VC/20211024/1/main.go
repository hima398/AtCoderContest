package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintSlice(s []int) {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintf(out, "%d", s[0])
	for i := 1; i < len(s); i++ {
		fmt.Fprintf(out, " %d", s[i])
	}
	fmt.Fprintln(out)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	m, k := nextInt(), nextInt()
	var ans []int
	if m == 1 {
		if k == 0 {
			ans = []int{0, 0, 1, 1}
			PrintSlice(ans)
			return
		} else {
			fmt.Println(-1)
			return
		}
	}
	if k >= 1<<m {
		fmt.Println(-1)
		return
	}
	for i := 0; i < 1<<m; i++ {
		if i != k {
			ans = append(ans, i)
		}
	}
	ans = append(ans, k)
	for i := (1 << m) - 1; i >= 0; i-- {
		if i != k {
			ans = append(ans, i)
		}
	}
	ans = append(ans, k)
	PrintSlice(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
