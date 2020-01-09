package gomock

/*
 * @author taohu
 * @date 2019/12/24
 * @time 上午11:21
 * @desc please describe function
**/

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {

}
