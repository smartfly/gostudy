[Go 1.9 sync.Map揭秘](https://colobu.com/2017/07/11/dive-into-sync-Map/)

[Go Performance Tools Cheat Sheet](https://steveazz.xyz/blog/go-performance-tools-cheat-sheet/)

# [go benchstat 工具](https://chenlujjj.github.io/go/benchstat/)

## 安装
```shell
go get -u golang.org/x/perf/cmd/benchstat
```
## 使用

```shell
$ benchstat -h
usage: benchstat [options] old.txt [new.txt] [more.txt ...]
options:
  -alpha α
    	consider change significant if p < α (default 0.05)
  -csv
    	print results in CSV form
  -delta-test test
    	significance test to apply to delta: utest, ttest, or none (default "utest")
  -geomean
    	print the geometric mean of each file
  -html
    	print results as an HTML table
  -norange
    	suppress range columns (CSV only)
  -sort order
    	sort by order: [-]delta, [-]name, none (default "none")
  -split labels
    	split benchmarks by labels (default "pkg,goos,goarch")
```

- 当输入是一个文件时：打印出该文件中的 benchmark 统计结果
- 当输入是一对文件时：分别打印出这两个文件的 benchmark 统计结果，以及它们的比较信息
- 当输入是三个及以上数量的文件时：分别打印出这多个文件的 benchmark 统计结果

## 举例

1. 执行两次benchmark测试, 并保存结果第一次
   
    ```shell
    ~/workspace/go/learning-go/map_api(master*) » go test -bench=. -test.benchmem -count=5 ./map_advance_test.go | tee -a old.txt
    goos: linux
    goarch: amd64
    BenchmarkMap-4              1573            727658 ns/op          687146 B/op        276 allocs/op
    BenchmarkMap-4              1627            730526 ns/op          687172 B/op        276 allocs/op
    BenchmarkMap-4              1627            750512 ns/op          687090 B/op        275 allocs/op
    BenchmarkMap-4              1522            726309 ns/op          687180 B/op        276 allocs/op
    BenchmarkMap-4              1543            718325 ns/op          687189 B/op        276 allocs/op
    BenchmarkCapMap-4           3679            310636 ns/op            2703 B/op          9 allocs/op
    BenchmarkCapMap-4           3415            310593 ns/op            2716 B/op          9 allocs/op
    BenchmarkCapMap-4           3504            310477 ns/op            2723 B/op          9 allocs/op
    BenchmarkCapMap-4           3682            310834 ns/op            2698 B/op          9 allocs/op
    BenchmarkCapMap-4           3850            310197 ns/op            2712 B/op          9 allocs/op
    PASS
    ok      command-line-arguments  16.665s
    
    ```
    第二次
   
   ```shell
   ~/workspace/go/learning-go/map_api(master*) » go test -bench=. -test.benchmem -count=5 ./map_advance_test.go | tee -a new.txt
   goos: linux
   goarch: amd64
   BenchmarkMap-4              1664            758519 ns/op          687294 B/op        276 allocs/op
   BenchmarkMap-4              1531            762922 ns/op          687149 B/op        275 allocs/op
   BenchmarkMap-4              1528            746472 ns/op          687178 B/op        276 allocs/op
   BenchmarkMap-4              1551            798821 ns/op          687170 B/op        275 allocs/op
   BenchmarkMap-4              1387            725624 ns/op          687181 B/op        275 allocs/op
   BenchmarkCapMap-4           3604            331148 ns/op            2731 B/op          9 allocs/op
   BenchmarkCapMap-4           3646            339677 ns/op            2725 B/op          9 allocs/op
   BenchmarkCapMap-4           3787            328241 ns/op            2726 B/op          9 allocs/op
   BenchmarkCapMap-4           3588            316845 ns/op            2710 B/op          9 allocs/op
   BenchmarkCapMap-4           3685            314174 ns/op            2735 B/op          9 allocs/op
   PASS
   
   ```
   
2. 查看单次benchmark统计结果

   ```shell
      ~/workspace/go/learning-go/map_api(master*) » benchstat old.txt
      name      time/op
      Map-4      731µs ± 3%
      CapMap-4   311µs ± 0%
      
      name      alloc/op
      Map-4      687kB ± 0%
      CapMap-4  2.71kB ± 0%
      
      name      allocs/op
      Map-4        276 ± 0%
      CapMap-4    9.00 ± 0%   
   ```

3. 比较两次的统计结果

   ```shell
   ~/workspace/go/learning-go/map_api(master*) » benchstat old.txt new.txt                                                                                                     smartfly@smartfly-virtual-machine
   name      old time/op    new time/op    delta
   Map-4        731µs ± 3%     758µs ± 5%    ~     (p=0.151 n=5+5)
   CapMap-4     311µs ± 0%     326µs ± 4%  +4.98%  (p=0.008 n=5+5)
   
   name      old alloc/op   new alloc/op   delta
   Map-4        687kB ± 0%     687kB ± 0%    ~     (p=0.548 n=5+5)
   CapMap-4    2.71kB ± 0%    2.73kB ± 1%    ~     (p=0.056 n=5+5)
   
   name      old allocs/op  new allocs/op  delta
   Map-4          276 ± 0%       275 ± 0%    ~     (p=0.095 n=4+5)
   CapMap-4      9.00 ± 0%      9.00 ± 0%    ~     (all equal)
   
   ```

   注意 delta 一列，值是~，没有显示百分比，说明两次测试结果没有显著差异。

## 参考资料

[Go性能工具小抄](https://github.com/gocn/translator/blob/master/2021/w23_go_performance_tools_cheat_sheet.md)