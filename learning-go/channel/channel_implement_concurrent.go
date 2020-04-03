package channel

import (
	"sync"
)

/*
 * @author taohu
 * @date 2020/4/2
 * @time 下午4:35
 * @desc 使用channel实现并发
**/

const batchValue = 10000

var wg sync.WaitGroup

// computer 计算
func computer(num []int, cnt int64, ch chan bool, resultCh chan int) {
	defer func() {
		wg.Done()
	}()
	ch <- true
	var sum int
	for _, val := range num {
		sum += val
	}
	resultCh <- sum
	<-ch
}

// currentComputer 并发计算
func currentComputer(numArray []int) int {
	var result int
	batchSize := batchValue
	totalCount := len(numArray)
	var times int
	if totalCount%batchSize == 0 {
		times = totalCount / batchSize
	} else {
		times = totalCount/batchSize + 1
	}
	cnt := int64(0)
	routineSizeCh := make(chan bool, 100)
	resultCh := make(chan int, times)
	for i := 0; i < times; i++ {
		startIndex := i * batchSize
		endIndex := (i + 1) * batchSize
		if endIndex > totalCount {
			endIndex = totalCount
		}
		wg.Add(1)
		go computer(numArray[startIndex:endIndex], cnt, routineSizeCh, resultCh)
	}
	wg.Wait()
	// 记得关闭channel, 不然容易死锁
	close(resultCh)
	for val := range resultCh {
		result += val
	}
	return result
}
