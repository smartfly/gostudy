# [Defer,Panic,and Recover](https://go.dev/blog/defer-panic-and-recover)

defer用于**清理释放资源**、**执行recover**和**释放锁**

The behavior of defer statements is straightforward and predictable. There are three simple rules:

1. A deferred function's arguments are evaluated when the defer statement is evaluated.
   
2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
   
3. Deferred functions may read and assign to the returning function's named return values.

## 参考资料
[语句系列之defer](https://studygolang.com/articles/12548?fr=sidebar)
[理解 Go defer](https://sanyuesha.com/2017/07/23/go-defer/)