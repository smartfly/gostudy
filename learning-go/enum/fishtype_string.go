// Code generated by "stringer -type=FishType"; DO NOT EDIT.

package enum

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[A-0]
	_ = x[B-1]
	_ = x[C-2]
	_ = x[D-3]
}

const _FishType_name = "ABCD"

var _FishType_index = [...]uint8{0, 1, 2, 3, 4}

func (i FishType) String() string {
	if i < 0 || i >= FishType(len(_FishType_index)-1) {
		return "FishType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FishType_name[_FishType_index[i]:_FishType_index[i+1]]
}
