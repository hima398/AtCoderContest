package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const size = 100

func PrintField(f [2*size+1][2*size+1]int) {
	for i:=0; i<len(f); i++ {
		fmt.Println(f[i])
	}
}

func Estimate() {
	var f [2*size+1][2*size+1]int
	for i:=0; i<=2*size; i++ {
		for j:=0; j<=2*size; j++ {
			f[i][j] = -1
		}
	}
	f[0][0] = 0
	f[0][1] = 0
	f[1][0] = 0

	var F func(x, y int) int
	F = func(x, y int) int {
		if f[x][y] >= 0 {
			return f[x][y]
		}

		// 次の手で負け確定を渡せる場合は勝ち
		canWin := false
		for i:=1; 2*i<=x; i++ {
			nx, ny := x-2*i, y+i
			if x!= nx || y != ny {
				canWin = canWin || F(nx, ny) == 0
			}
		}
		for i:=1; 2*i<=y; i++ {
			nx, ny := x+i, y-2*i
			if x!= nx || y != ny {
				canWin = canWin || F(nx, ny) == 0
			}
		}
		if canWin {
			f[x][y] = 1
		} else {
			f[x][y] = 0
		}
		return f[x][y]
	}
	for i:=0; i<=size; i++ {
		for j:=0; j<=size; j++ {
			//fmt.Println("i, j = ", i, j)
			F(i, j)
		}
	}
	//PrintField(f)
	type Pattern struct {
		x, y int
	}
	var alice, brown []Pattern
	for i:=0; i<=size; i++ {
		for j:=0; j<=size; j++ {
			if f[i][j] == 0 {
				brown = append(brown, Pattern{i, j})
			} else if f[i][j] == 1 {
				alice = append(alice, Pattern{i, j})
			}
		}
	}
	//fmt.Println(len(alice), len(brown))
	//fmt.Println(alice)
	fmt.Println(brown)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	//Estimate()

	x, y := nextInt(), nextInt()
	
	if Abs(x-y) <= 1 {
		fmt.Println("Brown")
	} else {
		fmt.Println("Alice")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
