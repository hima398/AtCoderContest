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
	if n == 1 {
		fmt.Println("Yes")
		fmt.Println("2")
		fmt.Println("1 1")
		fmt.Println("1 1")
		return
	}
	k := 1
	ok := false
	for k = 1; k <= n; k++ {
		if k*(k-1)/2 == n {
			ok = true
			break
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	if !ok {
		fmt.Fprintln(out, "No")
		return
	}

	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, k)
	m := make([][]int, k)
	sum := 1
	for i := k - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			//fmt.Fprintln(out, i, j)
			m[k-i-1] = append(m[k-i-1], sum+j)
			m[k-i+j] = append(m[k-i+j], sum+j)
		}
		sum += i
	}
	for i := 0; i < k; i++ {
		fmt.Fprintf(out, "%d ", len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			fmt.Fprintf(out, "%d ", m[i][j])
		}
		fmt.Fprintln(out, "")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
