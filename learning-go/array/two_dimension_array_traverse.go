package array

import "fmt"

func traverseByCommon(array [][]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("%v", array[i][j])
		}
		fmt.Println()
	}
}

func traverseByRange(array [][]int) {
	for i, v := range array {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v \t \n", i, j, v2)
		}
		fmt.Println()
	}

}
