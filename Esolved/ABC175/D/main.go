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

	const INF = int(1e18)
	//入力
	n, k := nextInt(), nextInt()
	p := nextIntSlice(n)
	c := nextIntSlice(n)

	for i := 0; i < n; i++ {
		// 0-indexedにする
		p[i]--
	}

	//マス目iから元の位置に戻るまでの巡回数
	ns := make([]int, n)
	//マス目iから元の位置に戻るまでに得られるスコア
	nc := make([]int, n)
	//1周期で得られる最大値
	nmxc := make([]int, n)
	for i := 0; i < n; i++ {
		pre, next := i, p[i] //e[i][0]
		ns[i]++
		nc[i] += c[next]
		nmxc[i] = nc[i]
		for next != i {
			pre = next
			next = p[pre] //e[pre][0]
			ns[i]++
			nc[i] += c[next]
			nmxc[i] = Max(nmxc[i], nc[i])
		}
	}
	//fmt.Println(ns)
	//fmt.Println(nc)
	//fmt.Println(nmxc)

	//マス目iからスタートして得られる最大スコア
	ms := make([]int, n)
	for i := 0; i < n; i++ {
		// k回移動する時に周回出来る数
		nr := k / ns[i]
		// nr回周回したあと動ける回数
		nk := k % ns[i]

		// nr回周回した後回収できるスコア
		sk := -INF
		idx := p[i]
		ss := 0
		if nk == 0 {
			sk = nmxc[i]
		} else {
			for j := 0; j < nk; j++ {
				ss += c[idx]
				sk = Max(sk, ss)
				idx = p[idx]
				//fmt.Println(i, idx, ss, sk)
			}
		}
		// 周回する度にスコアが増える場合
		if nc[i] > 0 {
			//1周以上周回出来る場合
			if nr > 0 {
				sr := (nr - 1) * nc[i]
				// sr-1周回って、最大を取るか、sr周回って残りを回収するかで比較
				if nk == 0 {
					ms[i] = Max(sr+nmxc[i], sr+nc[i])
				} else {
					ms[i] = Max(sr+nmxc[i], sr+nc[i]+Max(sk, 0))
				}
			} else {
				//スコアは増えるが周回できない場合
				if nk == 0 {
					ms[i] = nmxc[i]
				} else {
					ms[i] = sk
				}
			}
		} else {
			// 周回するほど損する場合は1周以内に取得できる最大値
			ms[i] = nmxc[i]
		}
	}
	//fmt.Println(ms)
	ans := -INF
	for _, v := range ms {
		//fmt.Println(i, v) //486, 2537
		ans = Max(ans, v)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
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
