package roman_to_integer

/*
 * @author taohu
 * @date 2020/5/5
 * @time 下午5:24
 * @desc please describe function
**/

func romanToIntAdvance(s string) int {
	value := 0
	if len(s) == 0 {
		return value
	}
	preNum := getValue(s[0:1])
	for i := 1; i < len(s); i++ {
		num := getValue(s[i : i+1])
		if preNum >= num {
			value += preNum
		} else {
			value -= preNum
		}
		preNum = num
	}
	value += preNum
	return value
}

func getValue(val string) int {
	var result int
	switch val {
	case "I":
		result = 1
	case "V":
		result = 5
	case "X":
		result = 10
	case "L":
		result = 50
	case "C":
		result = 100
	case "D":
		result = 500
	case "M":
		result = 1000
	}
	return result
}
