package str_int_utils

import "strconv"

/*
 * @author taohu
 * @date 19-5-21
 * @time 上午10:05
 * @desc string transfer int/int64
**/

// string转int
func Str2Int(strValue string) (int, error) {
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

// string转int64
func Str2Int64(strValue string) (int64, error) {
	int64Value, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64Value, nil
}
