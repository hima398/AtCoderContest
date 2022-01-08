package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pos struct {
	x, y int
}
type Vector [3]int

func NewVector(x, y int) *Vector {
	ret := new(Vector)
	ret[0] = x
	ret[1] = y
	ret[2] = 1.0
	return ret
}

type Matrix [3][3]int

func NewMatrix() *Matrix {
	return new(Matrix)
}

//時計回りに90度回転する行列を作る
//  0, 1, 0
// -1, 0, 0
//  0, 0, 1
func CreateRotate() (ret *Matrix) {
	ret = NewMatrix()
	ret[0][1] = 1
	ret[1][0] = -1
	ret[2][2] = 1
	return ret
}

//反時計回りに90度回転する行列を作る
// 0, -1, 0
// 1,  0, 0
// 0,  0, 1
func CreateReverseRotate() (ret *Matrix) {
	ret = NewMatrix()
	ret[0][1] = -1
	ret[1][0] = 1
	ret[2][2] = 1
	return ret
}

//x=pを軸を対称に反転する行列を作る
// -1, 0, 2p
//  0, 1, 0
//  0, 0, 1
func CreateXReflection(p int) (ret *Matrix) {
	ret = NewMatrix()
	ret[0][0] = -1
	ret[0][2] = 2 * p
	ret[1][1] = 1
	ret[2][2] = 1
	return ret
}

//y=pを軸を対称に反転する行列を作る
// 1,  0, 0
// 0, -1, 2p
// 0,  0, 1
func CreateYReflection(y int) (ret *Matrix) {
	ret = NewMatrix()
	ret[0][0] = 1
	ret[1][1] = -1
	ret[1][2] = 2 * y
	ret[2][2] = 1
	return ret
}

func (m *Matrix) mulVector(v *Vector) (ret *Vector) {
	ret = new(Vector)
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			ret[i] += m[i][k] * v[k]
		}
	}
	return ret
}

func (m *Matrix) mul(m2 *Matrix) (ret *Matrix) {
	ret = NewMatrix()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				ret[i][j] += m[i][k] * m2[k][j]
			}
		}
	}
	return ret
}

func PrintMatrix(m *Matrix) {
	for i := 0; i < 3; i++ {
		fmt.Println(m[i])
	}
}

func OpsOf(op, p int) *Matrix {
	switch op {
	case 1:
		return CreateRotate()
	case 2:
		return CreateReverseRotate()
	case 3:
		return CreateXReflection(p)
	case 4:
		return CreateYReflection(p)
	}
	return nil
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ps := make([]*Vector, n)
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		ps[i] = NewVector(x, y)
	}
	m := nextInt()
	ops := make([]*Matrix, m)
	for i := 0; i < m; i++ {
		op := nextInt()
		p := 0
		if op >= 3 {
			p = nextInt()
		}
		ops[i] = OpsOf(op, p)
	}
	for i := 1; i < m; i++ {
		ops[i] = ops[i].mul(ops[i-1])
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	q := nextInt()
	for i := 0; i < q; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--
		var ans *Vector
		if a < 0 {
			ans = ps[b]
		} else {
			ans = ops[a].mulVector(ps[b])
		}
		fmt.Fprintln(out, ans[0], ans[1])
	}

}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
