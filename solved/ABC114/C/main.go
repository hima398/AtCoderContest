package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	var a []int
	var F func(s string)
	F = func(s string) {
		if len(s) >= 10 {
			return
		}
		if strings.Contains(s, "3") && strings.Contains(s, "5") && strings.Contains(s, "7") {
			ns, _ := strconv.Atoi(s)
			a = append(a, ns)
		}
		F(s + "3")
		F(s + "5")
		F(s + "7")
	}
	F("")
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	//fmt.Println(len(a))
	n := nextInt()
	ans := 0
	for _, v := range a {
		if n >= v {
			//fmt.Println(v)
			ans++
		} else {
			break
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
