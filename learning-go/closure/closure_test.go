package closure

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
 * @author taohu
 * @date 2020/3/19
 * @time 下午8:15
 * @desc go 闭包demo
**/

// getSequence 不带参数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Test_getSequence(t *testing.T) {
	// nextNumber为一个函数, 函数i为0
	nextNumber := getSequence()
	// 调用nextNumber函数, i变量自增1并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// 创建新的函数nextNumber1并查看结果
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// add 带参数闭包使用
func add(a, b int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, a + b
	}
}

func Test_add(t *testing.T) {
	addFunc := add(1, 2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())
	fmt.Println(addFunc())
}

func GoAndWait(handlers ...func() error) (err error) {
	var wg sync.WaitGroup
	var once sync.Once
	for _, f := range handlers {
		wg.Add(1)
		go func(handler func() error) {

			defer func() {
				if e := recover(); e != nil {
					buf := make([]byte, 1024)
					buf = buf[:runtime.Stack(buf, false)]
					fmt.Errorf("[PANIC]%v\n%s\n", e, buf)
				}
				wg.Done()
			}()

			if e := handler(); e != nil {
				once.Do(func() {
					err = e
				})
			}
		}(f)
	}

	wg.Wait()

	return err
}

func TestGoroutineChannel(t *testing.T) {
	pIDs := []string{"1", "2", "3"}
	results := make(map[string]bool)
	resultsChan := make(chan map[string]bool, len(pIDs))
	var tasks []func() error
	for _, pID := range pIDs {
		tasks = append(tasks, func() error {
			resultsChan <- map[string]bool{pID: true}
			return nil
		})
	}
	_ = GoAndWait(tasks...)
	close(resultsChan)
	for val := range resultsChan {
		for k, v := range val {
			results[k] = v
		}
	}
	fmt.Println(results)
}

func TestGoroutineChannelClosure(t *testing.T) {
	pIDs := []string{"1", "2", "3"}
	results := make(map[string]bool)
	resultsChan := make(chan map[string]bool, len(pIDs))
	var tasks []func() error
	for _, pID := range pIDs {
		temp := pID //闭包
		tasks = append(tasks, func() error {
			resultsChan <- map[string]bool{temp: true}
			return nil
		})
	}
	_ = GoAndWait(tasks...)
	close(resultsChan)
	for val := range resultsChan {
		for k, v := range val {
			results[k] = v
		}
	}
	fmt.Println(results)
}
