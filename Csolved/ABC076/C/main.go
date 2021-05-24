package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func ReplaceQuestions(s, t string, idx int) (ret string) {
	ret = s[:idx] + t + s[idx+len(t):]
	return ret
}

func SolveCommentary(s, t string) (ans string) {
	ns, nt := len(s), len(t)
	idx := -1
	for i := ns - 1; i >= nt-1; i-- {
		if s[i] == '?' || s[i] == t[nt-1] {
			ok := true
			f := i - nt + 1
			for j := 0; j < nt; j++ {
				if s[f+j] == '?' || s[f+j] == t[j] {
					continue
				}
				ok = false
				break
			}
			if ok {
				idx = f
				break
			}
		}
	}
	if idx == -1 {
		return "UNRESTORABLE"
	}
	//fmt.Println(idx)
	ans = ReplaceQuestions(s, t, idx)
	//fmt.Println(ans)
	ans = strings.Replace(ans, "?", "a", -1)

	return ans
}

func FirstSolve(s, t string) string {
	var idxs []int
	for i := 0; i < len(s)-len(t)+1; i++ {
		if s[i] == '?' || s[i] == t[0] {
			ok := true
			for j := 0; j < len(t); j++ {
				//fmt.Println(string(sd[i+j]), string(t[j]))
				if s[i+j] == '?' || s[i+j] == t[j] {
					continue
				}
				ok = false
				break
			}
			if ok {
				idxs = append(idxs, i)
			}
		}
	}
	//fmt.Println(idxs)
	if len(idxs) == 0 {
		return "UNRESTORABLE"
	}
	var ans []string
	for _, idx := range idxs {
		ss := ReplaceQuestions(s, t, idx)
		//fmt.Println(ss)
		ss = strings.Replace(ss, "?", "a", -1)
		ans = append(ans, ss)
	}
	sort.Strings(ans)
	//fmt.Println(ans)
	return ans[0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	sd := nextString()
	t := nextString()

	//fmt.Println(FirstSolve(sd, t))
	fmt.Println(SolveCommentary(sd, t))
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
