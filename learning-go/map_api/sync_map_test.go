package map_api

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
 * @author taohu
 * @date 2020/2/21
 * @time 下午12:05
 * @desc please describe function
**/

func TestSyncMapAPI(t *testing.T) {
	// 开箱即用
	var sm sync.Map
	// store方法，添加元素
	sm.Store(1, "a")
	// load方法，获得value
	if v, ok := sm.Load(1); ok {
		fmt.Println(v)
	}
	//LoadOrStore方法，获取或者保存
	//参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	if vv, ok := sm.LoadOrStore(1, "c"); ok {
		fmt.Println(vv)
	}

	if v, ok := sm.Load(1); ok {
		fmt.Println(v)
	}

	if vv, ok := sm.LoadOrStore(2, "c"); !ok {
		fmt.Println(vv)
	}

	//遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
	sm.Range(func(k, v interface{}) bool {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})

	// 删除map
	sm.Delete(1)
	fmt.Println("删除后的结果")
	sm.Range(func(k, v interface{}) bool {
		sm.Delete(k)
		return true
	})

	sm.Range(func(k, v interface{}) bool {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
		return true
	})
}

func TestSyncMapDemo(t *testing.T) {
	var sm sync.Map
	go func() {
		t := time.Tick(time.Second * 2)
		a := 0
		for {
			select {
			case <-t:
				for i := a; i < a+2; i++ {
					sm.Store(i, i)
				}
				a++
			}
		}

	}()

	go func() {
		t := time.Tick(time.Second * 1)
		for {
			select {
			case <-t:
				cmdList := make([]int, 0)
				sm.Range(func(key, value interface{}) bool {
					fmt.Println(key)
					cmdList = append(cmdList, key.(int))
					sm.Delete(key)
					return true
				})
				fmt.Println("-------------------------")
			}
		}
	}()

	time.Sleep(time.Second * 120)
}
