## Golang sort的slice与sliceStable的区别
- ```func Slice(slice interface{}, less func(i, j int) bool)``` 不保证相同的元素的次序
- ```func SliceStable(slice interface{}, less func(i, j int) bool) ``` 保证排序后相同的元素的先后顺序不变