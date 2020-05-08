package zigzag_conversion

/*
 * @author taohu
 * @date 2020/5/4
 * @time 上午11:40
 * @desc Z字形转换
**/

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}

	record := make([]string, numRows)
	curRow := 0
	flag := -1
	for _, val := range s {
		record[curRow] = record[curRow] + string(val)
		if (curRow == 0) || (curRow == numRows-1) {
			flag = -flag
		}
		curRow += flag
	}

	ret := ""
	for _, val := range record {
		ret += val
	}

	return ret
}
