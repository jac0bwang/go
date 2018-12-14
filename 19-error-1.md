# 19 错误处理（上）

我们使用error类型的方式通常是，在函数声明的结果列表的最后，声明一个该类型的结果，同时在调用这个函数之后，先判断它返回的最后一个结果值是否“不为nil”。
如果这个值“不为nil”，那么就进入错误处理流程，否则就继续进行正常的流程。

卫述语句,前面讲函数用法的时候也提到过卫述语句。简单地讲，它就是被用来检查后续操作的前置条件并进行相应处理的语句。

在生成error类型值的时候用到了errors.New函数。这是一种最基本的生成错误值的方式。我们调用它的时候传入一个由字符串代表的错误信息，它会给返回给我们一个包含了这个错误信息的error类型值。该值的静态类型当然是error，而动态类型则是一个在errors包中的，包级私有的类型*errorString。

当我们想通过模板化的方式生成错误信息，并得到错误值时，可以使用fmt.Errorf函数。该函数所做的其实就是先调用fmt.Sprintf函数，得到确切的错误信息；再调用errors.New函数，得到包含该错误信息的error类型值，最后返回该值。

## 对于具体错误的判断，Go 语言中都有哪些惯用法？

由于error是一个接口类型，所以即使同为error类型的错误值，它们的实际类型也可能不同。这个问题还可以换一种问法，即：怎样判断一个错误值具体代表的是哪一类错误？

1. 对于类型在已知范围内的一系列错误值，一般使用类型断言表达式或类型switch语句来判断；
2. 对于已有相应变量且类型相同的一系列错误值，一般直接使用判等操作来判断；
3. 对于没有相应变量且类型未知的一系列错误值，只能使用其错误信息的字符串表示形式来做判断。

类型在已知范围内的错误值其实是最容易分辨的。就拿os包中的几个代表错误的类型os.PathError、os.LinkError、os.SyscallError和os/exec.Error来说，它们的指针类型都是error接口的实现类型，同时它们也都包含了一个名叫Err，类型为error接口类型的代表潜在错误的字段。

```go
func underlyingError(err error) error {
    switch err := err.(type) {
    case *os.PathError:
        return err.Err
    case *os.LinkError:
        return err.Err
    case *os.SyscallError:
        return err.Err
    case *exec.Error:
        return err.Err
    }
    return err
}
```

函数underlyingError的作用是：获取和返回已知的操作系统相关错误的潜在错误值。

只要类型不同，我们就可以如此分辨。但是在错误值类型相同的情况下，这些手段就无能为力了。在 Go 语言的标准库中也有不少以相同方式创建的同类型的错误值。

我们还拿os包来说，其中不少的错误值都是通过调用errors.New函数来初始化的，比如：os.ErrClosed、os.ErrInvalid以及os.ErrPermission，等等。

```go
err = underlyingError(err)
switch err {
case os.ErrClosed:
    fmt.Printf("error(closed)[%d]: %s\n", i, err)
case os.ErrInvalid:
    fmt.Printf("error(invalid)[%d]: %s\n", i, err)
case os.ErrPermission:
    fmt.Printf("error(permission)[%d]: %s\n", i, err)
}
```

好在我们总是能通过错误值的Error方法，拿到它的错误信息。其实os包中就有做这种判断的函数，比如：os.IsExist、os.IsNotExist和os.IsPermission

```go
err = underlyingError(err)
if os.IsExist(err) {
    fmt.Printf("error(exist)[%d]: %s\n", i, err)
} else if os.IsNotExist(err) {
    fmt.Printf("error(not exist)[%d]: %s\n", i, err)
} else if os.IsPermission(err) {
    fmt.Printf("error(permission)[%d]: %s\n", i, err)
} else {
    fmt.Printf("error(other)[%d]: %s\n", i, err)
}
```

## 经常用到或者看到的 3 个错误类型，它们所在的错误类型体系都是怎样的？你能画出一棵树来描述它们吗？

常用到的net和json包中的错误类型有：

1. AddrError
2. SyntaxError
3. MarshalerError