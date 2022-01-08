package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ans := make([][]string, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]string, n)
		for j := 0; j < n; j++ {
			ans[i][j] = "."
		}
	}
	paintRow := 1
	for k := 1; k < n; k++ {
		if paintRow == 1 {
			for j := 0; j < k; j++ {
				ans[k][j] = "#"
			}
		} else {
			for i := 0; i < k; i++ {
				ans[i][k] = "#"
			}
		}
		//行を黒く塗るか列を黒く塗るかを切り替える
		paintRow ^= 1
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, strings.Join(ans[i], ""))
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
