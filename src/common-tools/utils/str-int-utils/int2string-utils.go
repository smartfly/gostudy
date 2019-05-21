package str_int_utils

import "strconv"

/*
 * @author: taohu
 * @date: 19-5-21
 * @time: 上午10:15
 * @desc: int/int64 transfer to string
**/

// int转string
func Int2Str(intValue int) string {
	strValue := strconv.Itoa(intValue)
	return strValue
}

// int64转string
func Int642Str(int64Value int64) string {
	strValue := strconv.FormatInt(int64Value, 10)
	return strValue
}
