package roman_to_integer

/*
 * @author taohu
 * @date 2020/5/5
 * @time 下午12:06
 * @desc  罗马数字转整数
**/

func romanToInt(s string) int {

	reflectMap := map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40,
		"L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}

	value := 0
	for i := 0; i < len(s); {
		if i+1 < len(s) {
			a := s[i : i+2]
			if val, ok := reflectMap[a]; ok {
				value += val
				i = i + 2
			} else {
				a := s[i : i+1]
				value += reflectMap[a]
				i++
			}
		} else {
			a := s[i : i+1]
			value += reflectMap[a]
			i++
		}
	}
	return value
}
