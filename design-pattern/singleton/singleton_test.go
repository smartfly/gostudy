package singleton

import (
	"testing"
)

/*
 * @author taohu
 * @date 2020/1/16
 * @time 下午7:45
 * @desc please describe function
**/

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}
