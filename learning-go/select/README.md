# select
```select```只能应用于channel的操作，既可以用于channel的数据接收，也可以用于channeld的数据发送。

如果```select```的多个分支都满足条件，则会随机的选取其中一个满足条件分支。如语言规范中所说：
> If multiple cases can proceed, a uniform pseudo-random choice is made to decide which single communication will execute.

```case```语句的表达式可以为一个变量或者两个变量赋值

有```default```语句