package main

import "math/rand"

const MaxRand = 16 // 声明一个具名整形常量

/*
StatRandomNumbers 生成一些不大于MaxRand的非负随机整数，
并统计和返回小于和大于MaxRand/2的随机数个数。
输入参数numRands指定了要生成的随机数总数。
*/
func StatRandomNumbers(numRands int) (int, int) {
	var a, b int
	for i := 0; i < numRands; i++ {
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++ // 等价于: b = b+1
		}
	}
	return a, b
}

func main() {
	var num = 100
	x, y := StatRandomNumbers(100)
	print("Result: ", x, " + ", y, " = ", num, "? ")
	println(x+y == num)
}
