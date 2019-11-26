package area

/*
 * @author taohu
 * @date 2019/11/26
 * @time 下午7:28
 * @desc color
**/

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
