package mock

import (
	"github.com/golang/mock/gomock"
	gomock2 "go/uint-test/gomock"
	"testing"
)

/*
 * @author taohu
 * @date 2019/12/24
 * @time 上午11:29
 * @desc please describe function
**/

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(101).AnyTimes()

	m.
		EXPECT().
		Bar(gomock.Eq(101)).
		Return(103).
		AnyTimes()

	gomock2.SUT(m)
}
