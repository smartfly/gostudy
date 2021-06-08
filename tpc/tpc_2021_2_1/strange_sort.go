package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type InputScanner struct {
	Scanner *bufio.Scanner
}

func NewInputScanner() InputScanner {
	inputScanner := InputScanner{bufio.NewScanner(os.Stdin)}
	inputScanner.Scanner.Split(bufio.ScanWords)
	return inputScanner
}

func (input InputScanner) NextInt() int {
	input.Scanner.Scan()
	return str2int(input.Scanner.Text())
}

func str2int(s string) int {
	intVal, _ := strconv.Atoi(s)
	return intVal
}

func (input InputScanner) NextIntList(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		input.Scanner.Scan()
		s := input.Scanner.Text()
		result[i] = str2int(s)
	}
	return result
}

func main() {
	input := NewInputScanner()
	totalLine := input.NextInt()
	for i := 0; i < totalLine; i++ {
		num := input.NextInt()
		dataSlice := input.NextIntList(num)
		suc := sortInt(dataSlice)
		if suc {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func sortInt(dataSlice []int) bool {
	if len(dataSlice) <= 1 {
		return true
	}
	if len(dataSlice) == 2 && dataSlice[0] <= dataSlice[1] {
		return true
	} else if len(dataSlice) == 2 && dataSlice[0] > dataSlice[1] {
		return false
	}
	size := len(dataSlice)
	if size%2 != 0 {
		return true
	}
	intList1 := make([]int, 0, size/2)
	intList2 := make([]int, 0, size/2)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			intList1 = append(intList1, dataSlice[i])
		} else {
			intList2 = append(intList2, dataSlice[i])
		}
	}
	sort.Ints(intList1)
	sort.Ints(intList2)
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			result = append(result, intList1[i/2])
		} else {
			result = append(result, intList2[i/2])
		}
	}
	return sort.IntsAreSorted(result)
}
