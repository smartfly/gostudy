
# rand
golang 提供两种生成随机数方法：
- math/rand
- crypto/rand
其中 math/rand 提供伪随机数，crypto/rand 提供真正的随机数

## math/rand
- 不同种子可以得到相同的结果
- rand.New 来初始化的rand并不是并发安全的
- 建议直接使用 ```math.Intn()``` 并发安全，其实现全局共享一个锁

**切记**：
为了防止全局锁竞争问题，在使用```math/rand```的时候，首先都会想到自定义rand, 但是就容易整出来莫名其妙的问题。

### 为什么 math/rand 需要加锁呢？

大家都知道 math/rand 是伪随机, 但是在设置完seed后，rng.vec 数组的值基本上就确定下来了，这明显就不是随机了，为了增加随机性，通过 Uint64()
获取到随机数后，还会重新去设置rng.vec。由于存在并发获取随机数的需求，也就有了并发设置rng.vec的值，所以需要对rng.vec加锁保护。

### 为什么 ```rand.Intn()``` 有性能问题？
由于```rand.Intn()```使用全局锁，存在全局锁竞争问题。

[Go: 谨慎使用 math/rand 包中的默认随机数函数](https://mozillazg.com/2019/05/go-be-careful-of-math-rand-functions.html)
[一文完全掌握 Go math/rand](https://segmentfault.com/a/1190000039855767)
[golang实践-随机数的那点事儿](https://blog.csdn.net/qq_26981997/article/details/56837947)
[How can I generate a random int using the "crypto/rand" package?](https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package)