package string_to_integer_atoi

import (
	"fmt"
	"math"
)

/*
 * @author taohu
 * @date 2020/5/4
 * @time 上午11:50
 * @desc
**/

func myAtoi(str string) int {
	result := 0
	stat := "start"
	sign := 1

	table := map[string][4]string{
		"start":  [4]string{"start", "signed", "number", "end"},
		"signed": [4]string{"end", "end", "number", "end"},
		"number": [4]string{"end", "end", "number", "end"},
		"end":    [4]string{"end", "end", "end", "end"},
	}

	for _, val := range str {
		stat = table[stat][getCol(val)]
		fmt.Println(stat)
		if stat == "number" {
			result = result*10 + int(val-'0')
			if sign == -1 && result*sign < math.MinInt32 {
				return math.MinInt32
			} else if sign == 1 && result*sign > math.MaxInt32 {
				return math.MaxInt32
			}
		} else if stat == "signed" {
			switch val {
			case '+':
				sign = 1
			case '-':
				sign = -1
			}
		} else if stat == "end" {
			break
		}
	}

	value := sign * result

	return value

}

func getCol(c int32) int {
	if c == ' ' {
		return 0
	}

	if c == '+' || c == '-' {
		return 1
	}

	if c >= '0' && c <= '9' {
		return 2
	}

	return 3
}
