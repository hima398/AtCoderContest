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

	const p = 998244353
	const INF = 1001

	n, lk := nextInt(), nextInt()
	c := make([]string, lk)
	k := make([]int, lk)
	pat := make([]int, n)
	for i := 0; i < lk; i++ {
		c[i], k[i] = nextString(), nextInt()
		k[i]--
		if c[i] == "L" {
			pat[k[i]] = INF
			//k[i]より右側のカードはiの数字に書き換えられる
			for j := k[i] + 1; j < n; j++ {
				if pat[j] != INF {
					pat[j]++
				}
			}
		} else {
			// c[i] == "R"
			pat[k[i]] = INF
			// k[i]より左側はiの数字に書き換えられる
			for j := 0; j < k[i]; j++ {
				if pat[j] != INF {
					pat[j]++
				}
			}
		}
	}
	//fmt.Println(pat)
	ans := 1
	for i := 0; i < n; i++ {
		if pat[i] != INF {
			ans *= pat[i]
			ans %= p
		}
		//fmt.Println(ans)
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
