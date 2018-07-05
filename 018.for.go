/*
    for
    Go 只有一种循环结构 for 循环
    基本的 for 循环由三部分组成，用分号隔开：
        初始化语句: 在第一次迭代前执行
        条件表达式: 在每次迭代前求值
        后置语句: 在每次迭代的结尾执行
    
    初始化语句一般为一句短变量声明，该变量仅在 for 语句作用域可见。
    一旦条件表达式为 false, 循环结束
    注意： 与 c, java, javascript 不同, Go 的 for 语句后面没有小括号，大括号 {} 是必须的
*/

package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}