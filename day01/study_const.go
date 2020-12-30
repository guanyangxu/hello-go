// 常量相关知识

package main

import "fmt"

// 常量的定义格式: const constName [type] = value
// type 类型说明符可以省略，因为编译器可以根据值推断出类型
const PI = 3.1415926

// 并行赋值
//const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Monday = 1
	Tuesday = 2
	Wednesday, Thursday, Friday, Saturday = 3, 4, 5, 6
)

// 没有赋值的常量默认会应用上一行的赋值表达式
// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值就会自动加 1
const (
	a = iota // 0
	b        // 1
	c        // 2
	d = 5    // 5
	e        // 5
)

// 在遇到新的常量块或单个常量声明时，iota 就会被重置为 0，（就是每遇到 const 关键字，iota 就会重置为 0）
const f = iota

func main() {
	fmt.Println("PI 的值为:", PI)

	//fmt.Println("Tuesday 未赋值的值为:", Tuesday)

	fmt.Println("常量 a 的值为:", a)
	fmt.Println("常量 b 的值为:", b)
	fmt.Println("常量 d 的值为:", d)
	fmt.Println("常量 e 的值为:", e)

	fmt.Println("常量 f 的值为:", f)
}
