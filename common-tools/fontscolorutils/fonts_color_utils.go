package fontscolorutils

import "fmt"

/*
 * @author taohu
 * @date 2020/2/19
 * @time 下午5:47
 * @desc please describe function
**/

/*
	--------+-----------+-----------+
	|  前景	|	背景	|	颜色	|
	--------+-----------+-----------+
	|  30	|	40		|	黑色	|
	+-------+-----------+-----------+
	|  31	|	41		|	红色	|
	+-------+-----------+-----------+
	|  32	|	42		|	绿色	|
	+-------+-----------+-----------+
	|  33	|	43		|	黄色	|
	+-------+-----------+-----------+
	|  34	|	44		|	蓝色	|
	+-------+-----------+-----------+
	|  35	|	45		|	紫色	|
	+-------+-----------+-----------+
	|  36	|	46		|	青色	|
	+-------+-----------+-----------+
	|  37	|	47		|	白色	|
	+-------+-----------+-----------+

	+-------+-----------------------+
	|  代码	|	意义				|
	+-------+-----------------------+
	|  0	|	终端默认设置			|
	+-------+-----------------------+
	|  1	|	高亮显示				|
	+-------+-----------------------+
	|  4	|	使用下划线			|
	+-------+-----------------------+
	|  5	|	闪烁				|
	+-------+-----------------------+
	|  7	|	反白显示				|
	+-------+-----------------------+
	|  8	|	不可见				|
	+-------+-----------------------+
*/

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

const (
	BackgroundBlack  = iota + 40 // 黑色
	BackgroundRed                // 红色
	BackgroundGreen              // 绿色
	BackgroundYellow             // 黄色
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

const (
	DisplayDefault   = 0 // 终端默认设置
	DisplayHighLight = 1 // 高亮显示
	DisplayUnderline = 4 // 使用下划线
	DisplayFlash     = 5 // 闪烁
	DisplayReverse   = 7 // 反白显示
	DisplayInvisible = 8 // 不可见
)

func Black(msg string) string {
	return SetColor(msg, 0, 0, TextBlack)
}

func Red(msg string) string {
	return SetColor(msg, 0, 0, TextRed)
}

func Green(msg string) string {
	return SetColor(msg, 0, 0, TextGreen)
}

func Yellow(msg string) string {
	return SetColor(msg, 0, 0, TextYellow)
}

func Blue(msg string) string {
	return SetColor(msg, 0, 0, TextBlue)
}

func Magenta(msg string) string {
	return SetColor(msg, 0, 0, TextMagenta)
}

func Cyan(msg string) string {
	return SetColor(msg, 0, 0, TextCyan)
}

func White(msg string) string {
	return SetColor(msg, 0, 0, TextWhite)
}

func SetColor(msg string, displayType, bg, text int) string {
	return fmt.Sprintf(" %c[%d;%d;%dm%s%c[0m ", 0x1B, displayType, bg, text, msg, 0x1B)
}
