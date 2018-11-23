# 2 命令源码文件

源码文件又分为三种，即：命令源码文件、库源码文件和测试源码文件，它们都有着不同的用途和编写规则

这篇文章主要解析 **命令参数的接收和解析有关的一系列问题**

如果一个源码文件声明属于main包，并且包含一个**无参数声明**且**无结果声明**的main函数，那么它就是**命令源码文件**

```go
// File name: demo1.go
// 运行go run demo1.go命令后就会在屏幕（标准输出）中看到Hello, world!
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

## 注意

    当需要模块化编程时，我们往往会将代码拆分到多个文件，甚至拆分到不同的代码包中。但无论怎样，对于一个独立的程序来说，命令源码文件永远只会也只能有一个。如果有与命令源码文件同包的源码文件，那么它们也应该声明属于main包。

## 1. 命令源码文件怎样接收参数

首先，Go 语言标准库中有一个代码包专门用于接收和解析命令参数。这个代码包的名字叫flag

```go
// File name: demo2.go
package main

import (
    "flag" //[1]
    "fmt"
)

var name string
// 方式 2
//var name = flag.String("name", "everyone", "The greeting object.")

func init() {
    // 方式 1
    // 函数flag.StringVar接受 4 个参数。
    // 第 1 个参数是用于存储该命令参数的值的地址，具体到这里就是在前面声明的变量name的地址了，由表达式&name表示
    // 第 2 个参数是为了指定该命令参数的名称，这里是name。
    // 第 3 个参数是为了指定在未追加该命令参数时的默认值，这里是everyone。
    // 第 4 个函数参数，即是该命令参数的简短说明了，这在打印命令说明时会用到
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
    // 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量
    flag.Parse()
    fmt.Printf("Hello, %s!\n", name)
}
```

## 2. 怎样在运行命令源码文件的时候传入参数，又怎样查看参数的使用说明

运行如下命令就可以为参数name传值

```bash
$ go run demo2.go -name="Robert"
Hello, Robert!
$ go run demo2.go --help
Usage of /tmp/go-build233321288/b001/exe/demo2:
  -name string
        The greeting object. (default "everyone")
exit status 2
```

    /tmp/go-build233321288/b001/exe/demo2
这其实是go run命令构建上述命令源码文件时临时生成的可执行文件的完整路径

也可以先构建这个命令源码文件再运行生成的可执行文件

```bash
$ go build demo2.go
$ ./demo2 --help
Usage of ./demo2:
  -name string
        The greeting object. (default "everyone")
```

## 3. 怎样自定义命令源码文件的参数使用说明

这有很多种方式，最简单的一种方式就是对变量flag.Usage重新赋值。flag.Usage的类型是func()，即一种无参数声明且无结果声明的函数类型

flag.Usage变量在声明时就已经被赋值了，所以我们才能够在运行命令go run demo2.go --help时看到正确的结果

注意，对flag.Usage的赋值必须在调用flag.Parse函数之前

flag.ExitOnError的含义是，告诉命令参数容器，当命令后跟--help或者参数设置的不正确的时候，在打印命令参数使用说明后以状态码2结束当前程序

**状态码2**代表用户错误地使用了命令，而**flag.PanicOnError**与之的区别是在最后抛出**运行时恐慌（panic）**

全局的flag.CommandLine变量

私有的命令参数容器  cmdLine = flag.NewFlagSet

```go
// File name : demo3.go
package main

import (
        "flag"
        "fmt"
        "os"
)

var name string

// 方式3。私有的命令参数容器
//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
        // 方式2。 全局的flag.CommandLine变量
        flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
        //flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
        flag.CommandLine.Usage = func() {
                fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
                flag.PrintDefaults()
        }
        // 方式3。
        //cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
        flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
        // 方式1。
        //flag.Usage = func() {
        //      fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
        //      flag.PrintDefaults()
        //}
        // 方式3。
        //cmdLine.Parse(os.Args[1:])
        flag.Parse()
        fmt.Printf("Hello, %s!\n", name)
}
```

查看[更多关于 flag 的使用](https://golang.google.cn/pkg/flag/)，或者使用godoc命令在本地启动一个 [Go 语言文档服务器](使用godoc命令在本地启动一个 Go 语言文档服务器)

## 思考题

1. 默认情况下，我们可以让命令源码文件接受哪些类型的参数值？

    命令源码文件支持的参数:
    int(int|int64|uint|uint64),
    float(float|float64)
    string,
    bool,
    duration(时间),
    var(自定义)

2. 我们可以把自定义的数据类型作为参数值的类型吗？如果可以，怎样做？

    关键就是使用flag.var()，关键点在于需要实现flag包的Value接口