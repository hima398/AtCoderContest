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

	di := []int{-1, 0, 0, 1}
	dj := []int{0, -1, 1, 0}
	h, w := nextInt(), nextInt()
	f := make([][]bool, h)
	for i := 0; i < h; i++ {
		f[i] = make([]bool, w)
	}
	q := nextInt()
	uf := NewUnionFind(h, w)
	for i := 0; i < q; i++ {
		t := nextInt()
		if t == 1 {
			r, c := nextInt(), nextInt()
			r--
			c--
			f[r][c] = true
			for k := 0; k < 4; k++ {
				nr := r + di[k]
				nc := c + dj[k]
				if nr >= 0 && nr < h && nc >= 0 && nc < w && f[nr][nc] {
					uf.Unite(r, c, nr, nc)
				}
			}
		} else {
			// t == 2
			ra, ca, rb, cb := nextInt(), nextInt(), nextInt(), nextInt()
			ra--
			ca--
			rb--
			cb--
			if f[ra][ca] && f[rb][cb] && uf.ExistSameUnion(ra, ca, rb, cb) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type UnionFind struct {
	pari [][]int // parent numbers
	parj [][]int // parent numbers
	rank [][]int // height of tree
}

func NewUnionFind(h, w int) *UnionFind {
	if h <= 0 || w <= 0 {
		return nil
	}
	u := new(UnionFind)
	// for accessing index without minus 1
	u.pari = make([][]int, h)
	u.parj = make([][]int, h)
	u.rank = make([][]int, h)
	for i := 0; i < h; i++ {
		u.pari[i] = make([]int, w)
		u.parj[i] = make([]int, w)
		u.rank[i] = make([]int, w)
		for j := 0; j < w; j++ {
			u.pari[i][j] = i
			u.parj[i][j] = j
			u.rank[i][j] = 0
		}
	}
	return u
}

func (this *UnionFind) Find(i, j int) (int, int) {
	if this.pari[i][j] == i && this.parj[i][j] == j {
		return i, j
	} else {
		// compress path
		// ex. Find(4)
		// 1 - 2 - 3 - 4
		// 1 - 2
		//  L-3
		//  L 4

		this.pari[i][j], this.parj[i][j] = this.Find(this.pari[i][j], this.parj[i][j])
		return this.pari[i][j], this.parj[i][j]
	}
}

func (this *UnionFind) ExistSameUnion(i1, j1, i2, j2 int) bool {
	p1i, p1j := this.Find(i1, j1)
	p2i, p2j := this.Find(i2, j2)
	return p1i == p2i && p1j == p2j
}

//func (this *UnionFind) Unite(x, y int) {
func (this *UnionFind) Unite(i1, j1, i2, j2 int) {
	i1, j1 = this.Find(i1, j1)

	i2, j2 = this.Find(i2, j2)
	if i1 == i2 && j1 == j2 {
		return
	}
	// raink
	if this.rank[i1][j1] < this.rank[i2][j2] {
		this.pari[i1][j1] = i2
		this.parj[i1][j1] = j2
	} else {
		// this.rank[x] >= this.rank[y]
		this.pari[i2][j2], this.parj[i2][j2] = i1, j1
		if this.rank[i1][j1] == this.rank[i2][j2] {
			this.rank[i1][j1]++
		}
	}
}

func PrintUnionFind(u *UnionFind) {
	// for debuging. not optimize.
	fmt.Println(u.pari)
	fmt.Println(u.parj)
	fmt.Println(u.rank)
}
