package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve() {
	a, b, c := nextString(), nextInt(), nextInt()
	z := a[len(a)-1]
	if b == 0 {
		fmt.Println(1)
	}
	//bb := big.NewInt(6535897)
	//cc := big.NewInt(9323846)
	//bbcc := bb.Exp(bb, cc, big.NewInt(4))
	//fmt.Println(bbcc.String())
	if z == '0' || z == '1' || z == '5' || z == '6' {
		fmt.Println(string(z))
	} else if z == '4' || z == '9' {
		// B^Cが偶数か奇数か
		bc := Pow(b, c, 2)
		if z == '4' {
			if bc == 0 {
				fmt.Println(6)
			} else {
				fmt.Println(4)
			}
		} else {
			if bc == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(9)
			}
		}
	} else {
		bc := Pow(b, c, 4)
		if z == '2' {
			if bc == 0 {
				fmt.Println(6)
			} else if bc == 1 {
				fmt.Println(2)
			} else if bc == 2 {
				fmt.Println(4)
			} else {
				fmt.Println(8)
			}
		} else if z == '3' {
			if bc == 0 {
				fmt.Println(1)
			} else if bc == 1 {
				fmt.Println(3)
			} else if bc == 2 {
				fmt.Println(9)

			} else {
				fmt.Println(7)
			}
		} else if z == '7' {
			if bc == 0 {
				fmt.Println(1)
			} else if bc == 1 {
				fmt.Println(7)
			} else if bc == 2 {
				fmt.Println(9)

			} else {
				fmt.Println(3)
			}
		} else {
			if bc == 0 {
				fmt.Println(6)
			} else if bc == 1 {
				fmt.Println(8)
			} else if bc == 2 {
				fmt.Println(4)

			} else {
				fmt.Println(2)
			}
		}
	}
}

func Solve(a, b, c int) int {
	a %= 10
	ans := 0
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c := nextInt(), nextInt(), nextInt()
	fmt.Println(Solve(a, b, c))
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}
