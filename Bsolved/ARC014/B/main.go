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
	w := make([]string, n)
	for i := 0; i < n; i++ {
		w[i] = nextString()
	}
	m := make(map[string]bool)
	Ok := func(i int, s string) bool {
		// 前の終わりの文字から始まる
		ok := i == 0 || (i > 0 && s[0] == w[i-1][len(w[i-1])-1])
		// 同じ言葉を使っていない
		ok = ok && !m[s]
		// nで終わっていない
		// ok = w[len(w[i])-1] != 'n'
		return ok
	}
	IsOriginalTurn := 1
	for i := 0; i < n; i++ {
		if IsOriginalTurn == 1 {
			// 高橋くんのターン
			if !Ok(i, w[i]) {
				fmt.Println("LOSE")
				return
			}
		} else {
			// 高橋クンのターン
			if !Ok(i, w[i]) {
				fmt.Println("WIN")
				return
			}
		}
		//使った文字をメモ
		m[w[i]] = true
		// ターンの交代
		IsOriginalTurn ^= 1
	}
	fmt.Println("DRAW")
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
