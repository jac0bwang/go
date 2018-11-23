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
