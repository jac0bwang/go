# 32 context.Context类型

如果，我们不能在一开始就确定执行子任务的 goroutine 的数量，那么使用WaitGroup值来协调它们和分发子任务的 goroutine，就是有一定风险的。一个解决方案是：分批地启用执行子任务的 goroutine。

WaitGroup值是可以被复用的，但需要保证其计数周期的完整性。尤其是涉及对其Wait方法调用的时候，它的下一个计数周期必须要等到，与当前计数周期对应的那个Wait方法调用完成之后，才能够开始。

## 怎样使用context包中的程序实体，实现一对多的 goroutine 协作流程？

```go

func coordinateWithContext() {
    total := 12
    var num int32
    fmt.Printf("The number: %d [with context.Context]\n", num)
    cxt, cancelFunc := context.WithCancel(context.Background())
    for i := 1; i <= total; i++ {
        go addNum(&num, i, func() {
            if atomic.LoadInt32(&num) == int32(total) {
                cancelFunc()
            }
        })
    }
    <-cxt.Done()
    fmt.Println("End.")
}
```

在这个函数体中，我先后调用了context.Background函数和context.WithCancel函数，并得到了一个可撤销的context.Context类型的值（由变量cxt代表），以及一个context.CancelFunc类型的撤销函数（由变量cancelFunc代表）。

在后面那条唯一的for语句中，我在每次迭代中都通过一条go语句，异步地调用addNum函数，调用的总次数只依据了total变量的值。

请注意我给予addNum函数的最后一个参数值。它是一个匿名函数，其中只包含了一条if语句。这条if语句会“原子地”加载num变量的值，并判断它是否等于total变量的值。

如果两个值相等，那么就调用cancelFunc函数。其含义是，如果所有的addNum函数都执行完毕，那么就立即通知分发子任务的 goroutine。

这里分发子任务的 goroutine，即为执行coordinateWithContext函数的 goroutine。它在执行完for语句后，会立即调用cxt变量的Done函数，并试图针对该函数返回的通道，进行接收操作。

由于一旦cancelFunc函数被调用，针对该通道的接收操作就会马上结束，所以，这样做就可以实现“等待所有的addNum函数都执行完毕”的功能。

Context类型之所以受到了标准库中众多代码包的积极支持，主要是因为它是一种非常通用的同步工具。它的值不但可以被任意地扩散，而且还可以被用来传递额外的信息和信号。

由于Context类型实际上是一个接口类型，而context包中实现该接口的所有私有类型，都是基于某个数据类型的指针类型，所以，如此传播并不会影响该类型值的功能和安全。

## 问题 1：“可撤销的”在context包中代表着什么？“撤销”一个Context值又意味着什么？

这个接口中有两个方法与“撤销”息息相关。Done方法会返回一个元素类型为struct{}的接收通道。不过，这个接收通道的用途并不是传递元素值，而是让调用方去感知“撤销”当前Context值的那个信号。

## 问题 2：撤销信号是如何在上下文树中传播的？

context包中包含了四个用于繁衍Context值的函数。其中的WithCancel、WithDeadline和WithTimeout都是被用来基于给定的Context值产生可撤销的子值的。

context包的WithCancel函数在被调用后会产生两个结果值。第一个结果值就是那个可撤销的Context值，而第二个结果值则是用于触发撤销信号的函数。

在撤销函数被调用之后，对应的Context值会先关闭它内部的接收通道，也就是它的Done方法会返回的那个通道。

然后，它会向它的所有子值（或者说子节点）传达撤销信号。这些子值会如法炮制，把撤销信号继续传播下去。最后，这个Context值会断开它与其父值之间的关联。

我们通过调用context包的WithDeadline函数或者WithTimeout函数生成的Context值也是可撤销的。它们不但可以被手动撤销，还会依据在生成时被给定的过期时间，自动地进行定时撤销。这里定时撤销的功能是借助它们内部的计时器来实现的。

当过期时间到达时，这两种Context值的行为与Context值被手动撤销时的行为是几乎一致的，只不过前者会在最后停止并释放掉其内部的计时器。

最后要注意，通过调用context.WithValue函数得到的Context值是不可撤销的。撤销信号在被传播时，若遇到它们则会直接跨过，并试图将信号直接传给它们的子值。

## 问题 3：怎样通过Context值携带数据？怎样从中获取数据？

WithValue函数在产生新的Context值（以下简称含数据的Context值）的时候需要三个参数，即：父值、键和值。与“字典对于键的约束”类似，这里键的类型必须是可判等的。

原因很简单，当我们从中获取数据的时候，它需要根据给定的键来查找对应的值。不过，这种Context值并不是用字典来存储键和值的，后两者只是被简单地存储在前者的相应字段中而已。

Context类型的Value方法就是被用来获取数据的。在我们调用含数据的Context值的Value方法时，它会先判断给定的键，是否与当前值中存储的键相等，如果相等就把该值中存储的值直接返回，否则就到其父值中继续查找。

如果其父值中仍然未存储相等的键，那么该方法就会沿着上下文根节点的方向一路查找下去。

注意，除了含数据的Context值以外，其他几种Context值都是无法携带数据的。因此，Context值的Value方法在沿路查找的时候，会直接跨过那几种值。

如果我们调用的Value方法的所属值本身就是不含数据的，那么实际调用的就将会是其父辈或祖辈的Value方法。这是由于这几种Context值的实际类型，都属于结构体类型，并且它们都是通过“将其父值嵌入到自身”，来表达父子关系的。

提醒一下，Context接口并没有提供改变数据的方法。因此，在通常情况下，我们只能通过在上下文树中添加含数据的Context值来存储新的数据，或者通过撤销此种值的父值丢弃掉相应的数据。如果你存储在这里的数据可以从外部改变，那么必须自行保证安全。

## Context值在传达撤销信号的时候是广度优先的，还是深度优先的？其优势和劣势都是什么？

深度优先，看func (c *cancelCtx) cancel(removeFromParent bool, err error)方法的源代码。