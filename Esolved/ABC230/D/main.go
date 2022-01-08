package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Wall struct {
	l, r int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, d := nextInt(), nextInt()
	var ws []Wall
	for i := 0; i < n; i++ {
		w := Wall{nextInt(), nextInt()}
		ws = append(ws, w)
	}
	sort.Slice(ws, func(i, j int) bool {
		return ws[i].r < ws[j].r
	})
	var stack []Wall
	ans := 0
	for i := 0; i < n; i++ {
		// 壁をスタックに積んでおく
		if len(stack) == 0 || ws[i].l <= stack[0].r+d-1 {
			stack = append(stack, ws[i])
		} else {
			// スタックに保存していた壁と、ws[i]がまとめて壊せない
			ans++
			stack = []Wall{ws[i]}
		}
		//fmt.Println(stack)
	}
	if len(stack) > 0 {
		ans++
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
