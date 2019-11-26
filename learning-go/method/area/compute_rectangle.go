package area

/*
 * @author taohu
 * @date 2019/11/26
 * @time 下午7:24
 * @desc 计算长方形
**/

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}
