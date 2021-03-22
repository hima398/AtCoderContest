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

	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	var hans []int
	var ians []int
	inv := false

	for i := 0; i < n; i++ {
		if inv {
			ians = append(ians, a[i])
			inv = false
		} else {
			hans = append(hans, a[i])
			inv = true
		}
	}
	//fmt.Println(ans, inv)
	ans := make([]int, n)
	j := 0
	for i := len(ians) - 1; i >= 0; i-- {
		ans[j] = ians[i]
		j++
	}
	for i := 0; i < len(hans); i++ {
		ans[j] = hans[i]
		j++
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	if inv {
		fmt.Fprintf(out, "%d", ans[n-1])
		for i := n - 2; i >= 0; i-- {
			fmt.Fprintf(out, " %d", ans[i])
		}
		fmt.Fprintln(out, "")
	} else {
		fmt.Fprintf(out, "%d", ans[0])
		for i := 1; i < n; i++ {
			fmt.Fprintf(out, " %d", ans[i])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
