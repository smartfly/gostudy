# Go学习记录

## [Go语言之旅](https://tour.go-zh.org/welcome/1)

- basic syntax and data structures: 基本语法和数据结构
- methods and interfaces: 方法和接口
- Go's concurrency primitives: Go语言并发基础

## 开源框架 [gin-gonic/gin](https://github.com/gin-gonic/gin)

Use a Vendor tool like Govendor

1. go get govendor
    ```git
    go get github.com/kardianos/govendor
    ```
2. Create your project folder and cd inside
    ```git
    mkdir -p $GOPATH/src/github.com/myusername/project && cd "$_"
    ```
3. Vendor init your project and add gin
    ```git
    govendor init
    govendor fetch github.com/gin-gonic/gin@v1.3
    ```
4. Copy a starting template inside your project
    ```git
    curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
    ```
5. Run your project
    ```git
    go run main.go
    ```