package sync_wait_group

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2019/9/16
 * @time 上午11:23
 * @desc Golang sync.WaitGroup用法
**/

func TestGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	// 使用sleep为了等待goroutine都运行完毕
	time.Sleep(time.Second)
}

// 使用channel来实现等待多个goroutine完成任务
func TestGoroutineByChannel(t *testing.T) {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

// sync.WaitGroup实现
func TestGoroutineByWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	// 提前知道总数
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// sync.WaitGroup实现
func TestGoroutineByWaitGroupByUnknownNum(t *testing.T) {
	wg := sync.WaitGroup{}
	// 未知总数
	s1 := rand.NewSource(time.Now().Unix()) // 随机数种子
	r1 := rand.New(s1)                      // 创建rand
	randomNum := r1.Intn(100)               // 生成随机数
	fmt.Println("totalNum: ", randomNum)
	for i := 0; i < randomNum; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// waitGroup对象不是一个引用类型
func TestGoroutineByWaitGroupTrans(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go f(i, &wg)
	}
	wg.Wait()
}

func f(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}

// waitGroup对象不是一个引用类型
func TestGoroutineByWaitGroupRef(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go g(i, wg)
	}
	wg.Wait()
}

func g(i int, wg sync.WaitGroup) {
	fmt.Println(i)
	wg.Done()
}
