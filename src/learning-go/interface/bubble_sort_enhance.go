package _interface

/*
 * @author: taohu
 * @date: 19-6-27
 * @time: 上午9:39
 * @desc: 冒泡排序增强型
**/

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Xs []string

/*整数型实现*/
func (p Xi) Len() int {
	return len(p)
}

func (p Xi) Less(i int, j int) bool {
	return p[j] < p[i]
}

func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

/*字符串实现*/
func (p Xs) Len() int {
	return len(p)
}

func (p Xs) Less(i int, j int) bool {
	return p[j] < p[i]
}

func (p Xs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}
