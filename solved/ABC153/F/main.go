package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Monster struct {
	x, h int
}

type Bomb struct {
	l, r, d int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, d, a := nextInt(), nextInt(), nextInt()
	m := make([]Monster, n)
	for i := 0; i < n; i++ {
		m[i].x, m[i].h = nextInt(), nextInt()
	}
	sort.Slice(m, func(i, j int) bool { return m[i].x < m[j].x })
	ans := 0
	var q []Bomb
	dmg := 0
	for i := 0; i < n; i++ {
		for len(q) > 0 && q[0].r < m[i].x {
			dmg -= q[0].d
			//pop
			q = q[1:]
		}
		//i番目のモンスターの残りHP
		rmh := m[i].h - dmg
		if rmh <= 0 {
			continue
		}
		cnt := Ceil(m[i].h-dmg, a)
		ndmg := cnt * a
		dmg += ndmg
		ans += cnt
		q = append(q, Bomb{m[i].x, m[i].x + 2*d, ndmg})
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
