package mock_dynamic

import (
	"time"
)

/*
 * @author taohu
 * @date 2020/3/30
 * @time 下午3:29
 * @desc please describe function
**/

var nowFn = time.Now

func GetTime24Inner(createAt time.Time) bool {
	now := nowFn()
	duration := now.Sub(createAt)
	if duration > 24*time.Hour {
		return true
	}
	return false
}
