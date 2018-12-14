# 14 接口类型的合理运用

在 Go 语言的语境中，当我们在谈论“接口”的时候，一定指的是接口类型。因为接口类型与其他数据类型不同，它是没法被值化的，或者说是没法被实例化的。

我们既不能通过调用new函数或make函数创建出一个接口类型的值，也无法用字面量来表示一个接口类型的值。对于某一个接口类型来说，如果没有任何数据类型可以作为它的实现，那么该接口的值就不可能存在。

通过关键字type和interface，我们可以声明出接口类型。接口类型的类型字面量与结构体类型的看起来有些相似，它们都用花括号包裹一些核心信息。只不过，结构体类型包裹的是它的字段声明，而接口类型包裹的是它的方法定义。

注意，接口类型声明中的这些方法所代表的就是该接口的方法集合。一个接口的方法集合就是它的全部特征。

怎样判定一个数据类型的某一个方法实现的就是某个接口类型中的某个方法呢？

有两个充分必要条件，一个是“两个方法的签名需要完全一致”，另一个是“两个方法的名称要一模一样”。显然，这比判断一个函数是否实现了某个函数类型要更加严格一些。

```go
dog := Dog{"little pig"}
var pet Pet = &dog
```

这里有几个名词需要你先记住。对于一个接口类型的变量来说，例如上面的变量pet，我们赋给它的值可以被叫做它的实际值（也称动态值），而该值的类型可以被叫做这个变量的实际类型（也称动态类型）。

比如，我们把取址表达式&dog的结果值赋给了变量pet，这时这个结果值就是变量pet的动态值，而此结果值的类型*Dog就是该变量的动态类型。

动态类型这个叫法是相对于静态类型而言的。对于变量pet来讲，它的静态类型就是Pet，并且永远是Pet，但是它的动态类型却会随着我们赋给它的动态值而变化。

在我们给一个接口类型的变量赋予实际的值之前，它的动态类型是不存在的。

你需要想办法搞清楚接口类型的变量（以下简称接口变量）的动态值、动态类型和静态类型都是什么意思。因为我会在后面基于这些概念讲解更深层次的知识。

## 当我们为一个接口变量赋值时会发生什么？

有一条通用的规则需要你知晓：如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那个值，而是该值的一个副本。

接口类型值的存储方式和结构:

接口类型本身是无法被值化的。在我们赋予它实际的值之前，它的值一定会是nil，这也是它的零值。

反过来讲，一旦它被赋予了某个实现类型的值，它的值就不再是nil了。不过要注意，即使我们像前面那样把dog的值赋给了pet，pet的值与dog的值也是不同的。这不仅仅是副本与原值的那种不同。

当我们给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构中。

严格来讲，这样一个变量的值其实是这个专用数据结构的一个实例，而不是我们赋给该变量的那个实际的值。所以我才说，pet的值与dog的值肯定是不同的，无论是从它们存储的内容，还是存储的结构上来看都是如此。不过，我们可以认为，这时pet的值中包含了dog值的副本。

我们就把这个专用的数据结构叫做iface吧，在 Go 语言的runtime包中它其实就叫这个名字。

iface的实例会包含两个指针，一个是指向类型信息的指针，另一个是指向动态值的指针。这里的类型信息是由另一个专用数据结构的实例承载的，其中包含了动态值的类型，以及使它实现了接口的方法和调用它们的途径，等等。

总之，接口变量被赋予动态值的时候，存储的是包含了这个动态值的副本的一个结构更加复杂的值。

## 问题 1：接口变量的值在什么情况下才真正为nil？

虽然被包装的动态值是nil，但是pet的值却不会是nil，因为这个动态值只是pet值的一部分而已。

顺便说一句，这时的pet的动态类型就存在了，是*Dog。我们可以通过fmt.Printf函数和占位符%T来验证这一点，另外reflect包的TypeOf函数也可以起到类似的作用。

我们把nil赋给了pet，但是pet的值却不是nil。

在 Go 语言中，我们把由字面量nil表示的值叫做无类型的nil。这是真正的nil，因为它的类型也是nil的。虽然dog2的值是真正的nil，但是当我们把这个变量赋给pet的时候，Go 语言会把它的类型和值放在一起考虑。

只要我们把一个有类型的nil赋给接口变量，那么这个变量的值就一定不会是那个真正的nil。因此，当我们使用判等符号==判断pet是否与字面量nil相等的时候，答案一定会是false。

那么，怎样才能让一个接口变量的值真正为nil呢？要么只声明它但不做初始化，要么直接把字面量nil赋给它。

## 问题 2：怎样实现接口之间的组合？

接口类型间的嵌入也被称为接口的组合。我在前面讲过结构体类型的嵌入字段，这其实就是在说结构体类型间的嵌入。

接口类型间的嵌入要更简单一些，因为它不会涉及方法间的“屏蔽”。只要组合的接口之间有同名的方法就会产生冲突，从而无法通过编译，即使同名方法的签名彼此不同也会是如此。因此，接口的组合根本不可能导致“屏蔽”现象的出现。

与结构体类型间的嵌入很相似，我们只要把一个接口类型的名称直接写到另一个接口类型的成员列表中就可以了。

```go
type Animal interface {
    ScientificName() string
    Category() string
}

type Pet interface {
    Animal
    Name() string
}
```

Go 语言团队鼓励我们声明体量较小的接口，并建议我们通过这种接口间的组合来扩展程序、增加程序的灵活性。

接口变量的值并不等同于这个可被称为动态值的副本。它会包含两个指针，一个指针指向动态值，一个指针指向类型信息。

基于此，即使我们把一个值为nil的某个实现类型的变量赋给了接口变量，后者的值也不可能是真正的nil。虽然这时它的动态值会为nil，但它的动态类型确是存在的。

请记住，除非我们只声明而不初始化，或者显式地赋给它nil，否则接口变量的值就不会为nil。

## 思考题

如果我们把一个值为nil的某个实现类型的变量赋给了接口变量，那么在这个接口变量上仍然可以调用该接口的方法吗？如果可以，有哪些注意事项？如果不可以，原因是什么？
