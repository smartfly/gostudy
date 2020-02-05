package singleton

import "sync"

/*
 * @author taohu
 * @date 2020/1/16
 * @time 下午7:43
 * @desc 单例模式 使用懒惰模式的单例模式, 使用双重检查加锁保证线程安全
**/

//Singleton 单例模式类
type Singleton struct{}

// 内部实现不对外暴露
var singleton *Singleton
var once sync.Once

// GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
