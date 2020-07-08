package baisc

/*
 * @author taohu
 * @date 2020/6/17
 * @time 上午10:26
 * @desc please describe function
**/

var bitSetValues []uint32

func createBitSet(size uint32) {
	bitSetValues = make([]uint32, (size>>5)+1) // divide by 32
}

func get(pos uint32) bool {
	wordNumber := pos >> 5  // divide by 32
	bitNumber := pos & 0x1F // mod 32
	return bitSetValues[wordNumber]&(1<<bitNumber) != 0
}

func set(pos int) {
	wordNumber := pos >> 5  // divide by 32
	bitNumber := pos & 0x1F // mod 32
	bitSetValues[wordNumber] = 1 << bitNumber
}
