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
		var a, b []int
		a = myScanner.NextIntArray(n)
		b = myScanner.NextIntArray(n)
		succ := cal(a, b, k)
		if succ {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func cal(a []int, b []int, k int) bool {
	teamNum := len(a)
	tencentMax := a[0] + b[0]
	biggerNum := 0
	for i := 1; i < teamNum; i++ {
		if a[i] > tencentMax {
			biggerNum++
		}
	}
	return biggerNum < k
}
