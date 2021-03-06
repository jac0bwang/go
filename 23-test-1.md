# 23 测试的基本规则和流程 （上）

对于程序或软件的测试也分很多种，比如：单元测试、API 测试、集成测试、灰度测试，等等。本模块会主要针对单元测试进行讲解。

单元测试，它又称程序员测试。顾名思义，这就是程序员们本该做的自我检查工作之一。

可以为 Go 程序编写三类测试，即：功能测试（test）、基准测试（benchmark，也称性能测试），以及示例测试（example）。

一般情况下，一个测试源码文件只会针对于某个命令源码文件，或库源码文件（以下简称被测源码文件）做测试，所以我们总会（并且应该）把它们放在同一个代码包内。

测试源码文件的主名称应该以被测源码文件的主名称为前导，并且必须以“_test”为后缀。例如，如果被测源码文件的名称为 demo52.go，那么针对它的测试源码文件的名称就应该是 demo52_test.go。

每个测试源码文件都必须至少包含一个测试函数。并且，从语法上讲，每个测试源码文件中，都可以包含用来做任何一类测试的测试函数，即使把这三类测试函数都塞进去也没有问题。我通常就是这么做的，只要把控好测试函数的分组和数量就可以了。

## Go 语言对测试函数的名称和签名都有哪些规定？

下面三个内容。

1. 对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明。
2. 对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是*testing.B类型的。
3. 对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定。

这个问题的目的一般有两个。

第一个目的当然是考察 Go 程序测试的基本规则。

第二个目的是作为一个引子，引出第二个问题，即：go test命令执行的主要测试流程是什么？

只有测试源码文件的名称对了，测试函数的名称和签名也对了，当我们运行go test命令的时候，其中的测试代码才有可能被运行。

go test命令在开始运行时，会先做一些准备工作，比如，确定内部需要用到的命令，检查我们指定的代码包或源码文件的有效性，以及判断我们给予的标记是否合法，等等。在准备工作顺利完成之后，go test命令就会针对每个被测代码包，依次地进行构建、执行包中符合要求的测试函数，清理临时文件，打印测试结果。这就是通常情况下的主要测试流程。

请注意上述的“依次”二字。对于每个被测代码包，go test命令会串行地执行测试流程中的每个步骤。

为了加快测试速度，它通常会并发地对多个被测代码包进行功能测试，只不过，在最后打印测试结果的时候，它会依照我们给定的顺序逐个进行，这会让我们感觉到它是在完全串行地执行测试流程。

另一方面，由于并发的测试会让性能测试的结果存在偏差，所以性能测试一般都是串行进行的。更具体地说，只有在所有构建步骤都做完之后，go test命令才会真正地开始进行性能测试。

并且，下一个代码包性能测试的进行，总会等到上一个代码包性能测试的结果打印完成才会开始，而且性能测试函数的执行也都会是串行的。

## 思考题

你还知道或用过testing.T类型和testing.B类型的哪些方法？它们都是做什么用的？

    testing.T 的部分功能有（判定失败接口，打印信息接口）
    testing.B 拥有testing.T 的全部接口，同时还可以统计内存消耗，指定并行数目和操作计时器等。