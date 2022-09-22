# Are all runtime errors recoverable in Go?

## 错误类型

- error
- panic
- throw

## 致命场景

- 并发读写map
- 堆栈内存耗尽
- 将nil函数作为goroutine启动
- goroutines死锁
- 线程限制耗尽
- 超时可用内存(OOM)

[面试官：程序崩溃一定能recover吗？有无法恢复的致命场景？](https://mp.weixin.qq.com/s?__biz=Mzg5NDY2MDk4Mw==&mid=2247487663&idx=1&sn=e42b4f65e5e0d8c443e29e775063c7a0&chksm=c01d68b7f76ae1a1413924c0acc18d1322ddc95aa39dad8eea9846a23efc028d200f6278a3bc&mpshare=1&scene=1&srcid=1230Orm5zwoaWESIb4ldiu2c&sharer_sharetime=1640841161646&sharer_shareid=e365d5847184bcd78a34746e5d45525a&version=3.1.23.90391&platform=mac#rd)