# golang 单元测试

Go自带了测试框架和工具，在testing包中，以便完成单元测试(T类型)和性能测试(B类型)。
一般测试代码放在*_test.go文件中，与被测试代码放于同一个包中。

## uint test
测试函数名称格式是：Test[^a-zA-z],即以Test开头，跟上非小写字符开头的字符串。
每个测试函数都接受一个*testing.T类型参数，用于输出信息或者中断测试。

测试方法有：
- Fail: 标记失败，但继续执行当前测试函数
- FailNow: 失败，立即终止当前测试函数执行
- Log: 输出错误信息
- Error: Fail + Log
- Fatal: FailNow + Log
- Skip: 跳过当前函数，通常用于未完成的测试用例

运行```go test```命令，自动搜集所有的测试文件(*_test.go)，提取全部测试函数。输出结果包括：出错的测试函数名称、执行时长和错误信息

```shell
smartfly@smartfly-virtual-machine:~/workspace/go/uint-test$ go test
--- FAIL: TestAdd2 (0.00s)
    add_test.go:26: result is wrong!
FAIL
exit status 1
FAIL    go/uint-test    0.002s

```

go test还有以下参数：

- -v: 显示所有测试函数运行细节

  ```shell
  go test -v
  ```

- -run regex: 指定要执行的测试函数

  ```shell
  go test -run 函数名
  ```

## performance test

性能测试函数以Benchmark开头，参数类型*testing.B，可与Test函数放在同一个文件中。默认情况下，go test不执行Benchmark测试，必须用“-bench <pattern>”指定性能测试函数。

B类型有以下参数：

- benchmem: 输出内存分配统计
- benchtime: 指定测试时间
- cpu: 指定GOMAXPROCS
- timeout: 超时限制

# 参考资料

[Go单元测试&性能测试](https://www.jianshu.com/p/1adc69468b6f)