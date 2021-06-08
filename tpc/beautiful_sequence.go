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
		data := myScanner.NextIntArray(n)
		succ := calc(data)
		if succ {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func calc(data []int) bool {
	var a, b, c int
	for _, v := range data {
		if v > 0 {
			a++
		} else if v < 0 {
			b++
		} else {
			c++
		}
	}
	if diff(a, b) > c+1 {
		return false
	}
	return true
}
func diff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
