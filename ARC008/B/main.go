package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	_ = nextInt()
	_ = nextInt()
	name := nextString()
	kit := nextString()
	mn, mk := make(map[rune]int), make(map[rune]int)
	for _, v := range name {
		mn[v]++
	}
	for _, v := range kit {
		mk[v]++
	}
	var answers []int
	for i := 0; i < 26; i++ {
		v := rune('A' + i)
		//必要な文字がキットにない
		if mn[v] > 0 && mk[v] == 0 {
			fmt.Println(-1)
			return
		}
		if mn[v] > mk[v] {
			answers = append(answers, Ceil(mn[v], mk[v]))
		}
	}
	if len(answers) == 0 {
		fmt.Println(1)
		return
	}
	sort.Ints(answers)
	fmt.Println(answers[len(answers)-1])
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Ceil(x, y int) int {
	return (x + y - 1) / y
}
