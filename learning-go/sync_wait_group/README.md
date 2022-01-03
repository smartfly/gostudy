# WaitGroup用法

# WaitGroup注意事项
- WaitGroup对象不是一个引用类型，在通过函数传值得时候需要使用地址
    ```golang
    func main() {
        wg := sync.WaitGroup{}  //所以这句可以换成 var wg sync.WaitGroup；
        wg.Add(100)
        for i := 0; i < 100; i++ {
            go f(i, &wg)
        }
        wg.Wait()
    }
    // 一定要通过指针传值，不然进程会进入死锁状态
    func f(i int, wg *sync.WaitGroup) {
        fmt.Println(i)
        wg.Done()
    }
    ```
- 计数器不能为负值，在使用```Add()```不能给wg设置为一个负值，否则代码将会报错

[Golang WaitGroup 原理深度剖析](https://www.cyhone.com/articles/golang-waitgroup/)

[Golang package sync 剖析(二)： sync.WaitGroup](https://segmentfault.com/a/1190000021653777)