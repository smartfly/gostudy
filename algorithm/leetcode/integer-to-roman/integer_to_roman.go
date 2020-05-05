package integer_to_roman

import "strings"

/*
 * @author taohu
 * @date 2020/5/5
 * @time 上午11:17
 * @desc 整数转罗马数字
**/

func intToRoman(num int) string {

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	strList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	builder := strings.Builder{}

	for i, v := range values {
		for num >= v {
			num = num - v
			builder.WriteString(strList[i])
		}
	}
	return builder.String()
}
