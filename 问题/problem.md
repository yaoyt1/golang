## 问题

**问题 9.1**

a）一个包能分成多个源文件么？

b）一个源文件是否能包含多个包？

**练习 9.3**

创建一个程序 main_greetings.go 能够和用户说 "Good Day" 或者 "Good Night"。不同的问候应该放到单独的 greetings 包中。

在同一个包中创建一个 `IsAM` 函数返回一个布尔值用来判断当前时间是 AM 还是 PM，同样创建 `IsAfternoon` 和 `IsEvening` 函数。

使用 main_greetings 作出合适的问候(提示：使用 time 包)。

**练习 9.4** 创建一个程序 main_oddven.go 判断前 100 个整数是不是偶数，将判断所用的函数编写在 even 包里。

**练习 9.5** 使用第 6.6 节的斐波那契程序：

1）将斐波那契功能放入自己的 fibo 包中并通过主程序调用它，存储最后输入的值在函数的全局变量。

2）扩展 fibo 包将通过调用斐波那契的时候，操作也作为一个参数。实验 "+" 和 “*”

main_fibo.go / fibonacci.go

**总结**

在 Go 中，类型就是类（数据和关联的方法）。Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。

在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 **组件编程（Component Programming）**。

许多开发者说相比于类继承，Go 的接口提供了更强大、却更简单的多态行为。

**备注**

如果真的需要更多面向对象的能力，看一下 [`goop`](https://github.com/losalamos/goop) 包（Go Object-Oriented Programming），它由 Scott Pakin 编写: 它给 Go 提供了 JavaScript 风格的对象（基于原型的对象），并且支持多重继承和类型独立分派，通过它可以实现你喜欢的其他编程语言里的一些结构。

**问题 10.1**

我们在某个类型的变量上使用点号调用一个方法：`variable.method()`，在使用 Go 以前，在哪儿碰到过面向对象的点号？

**问题 10.2**

a）假设定义： `type Integer int`，完成 `get()` 方法的方法体: `func (p Integer) get() int { ... }`。

b）定义： `func f(i int) {}; var v Integer` ，如何就 v 作为参数调用f？

c）假设 `Integer` 定义为 `type Integer struct {n int}`，完成 `get()` 方法的方法体：`func (p Integer) get() int { ... }`。

d）对于新定义的 `Integer`，和 b）中同样的问题。

**练习 10.13** celsius.go

为 float64 定义一个别名类型 `Celsius`，并给它定义 `String()`，它输出一个十进制数和 °C 表示的温度值。

**练习 10.14** days.go

为 int 定义一个别名类型 `Day`，定义一个字符串数组它包含一周七天的名字，为类型 `Day` 定义 `String()` 方法，它输出星期几的名字。使用 `iota` 定义一个枚举常量用于表示一周的中每天（MO、TU...）。

**练习 10.15** timezones.go

为 int 定义别名类型 `TZ`，定义一些常量表示时区，比如 UTC，定义一个 map，它将时区的缩写映射为它的全称，比如：`UTC -> "Universal Greenwich time"`。为类型 `TZ` 定义 `String()` 方法，它输出时区的全称。

**练习 10.16** stack_arr.go/stack_struct.go

实现栈（stack）数据结构：

![](images/10.7_fig.jpg?raw=true)

它的格子包含数据，比如整数 i、j、k 和 l 等等，格子从底部（索引 0）至顶部（索引 n）来索引。这个例子中假定 `n=3`，那么一共有 4 个格子。

一个新栈中所有格子的值都是 0。

将一个新值放到栈的最顶部一个空（包括零）的格子中，这叫做`push`。

获取栈的最顶部一个非空（非零）的格子的值，这叫做`pop`。
现在可以理解为什么栈是一个后进先出（LIFO）的结构了吧。

为栈定义一个`Stack` 类型，并为它定义 `Push` 和 `Pop` 方法，再为它定义 `String()` 方法（用于调试）输出栈的内容，比如：`[0:i] [1:j] [2:k] [3:l]`。

1）stack_arr.go：使用长度为 4 的 int 数组作为底层数据结构。

2）stack_struct.go：使用包含一个索引和一个 int 数组的结构体作为底层数据结构，索引表示第一个空闲的位置。

3）使用常量 LIMIT 代替上面表示元素个数的 4 重新实现上面的 1）和 2），使它们更具有一般性。