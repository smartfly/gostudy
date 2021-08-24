package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type MyScanner struct {
	Scanner *bufio.Scanner
}

func NewMyScanner() MyScanner {
	myScanner := MyScanner{bufio.NewScanner(os.Stdin)}
	myScanner.Scanner.Split(bufio.ScanWords)
	return myScanner
}

func (myScanner MyScanner) NextInt() int {
	myScanner.Scanner.Scan()
	return String2int(myScanner.Scanner.Text())
}

func String2int(s string) int {
	t, _ := strconv.Atoi(s)
	return t
}

func (myScanner MyScanner) NextIntArray(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		myScanner.Scanner.Scan()
		s := myScanner.Scanner.Text()
		result[i] = String2int(s)
	}
	return result
}

func main() {
	myScanner := NewMyScanner()
	T := myScanner.NextInt()
	for i := 0; i < T; i++ {
		n := myScanner.NextInt()
		k := myScanner.NextInt()
		a := myScanner.NextIntArray(n)
		roundNum, winNum, succ := cal(n, k, a)
		if succ {
			fmt.Println(roundNum, winNum)
		} else {
			fmt.Println("INF")
		}
	}
}

func cal(n, k int, a []int) (roundNum, winNum int, suc bool) {
	exist := make(map[int]struct{})
	for _, val := range a {
		exist[val] = struct{}{}
	}
	if len(exist) == 1 {
		return k, a[0], true
	}
	var curWin int
	for i := 0; i < n*2; i++ {
		if roundNum >= k {
			return i, curWin, true
		}
		winner := a[i%n]
		if curWin == winner {
			roundNum++
		} else {
			curWin = winner
			roundNum = 1
		}
	}
	return 0, 0, false
}
