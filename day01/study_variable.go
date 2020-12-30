// 变量的相关知识
// 一个变量（常量、类型、函数）在程序中都有一定的作用范围，称之为作用域，如果一个变量在函数体外声明，则被认为是全局变量，可以在
// 整个包甚至外部包（被导出后）使用，不管你声明在哪个源文件里或在哪个源文件里调用该变量
//
// 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量

package main

import (
	"fmt"
	"os"
	"runtime"
)

// 声明形式：var variableName type
//var var_a int
//var var_b bool
//var str string

// 默认值
// int: 0, float: 0.0, bool: false, string: "", 指针: nil
var (
	var_a int = 10
	var_b bool
	str string
	ff int64
)

func init() {
	fmt.Println("这个是 init 函数")
}

func main()  {
	goos := runtime.GOOS
	fmt.Printf("这个操作系统是：%s\n", goos)

	var path string = os.Getenv("PATH")
	fmt.Println("路径是:", &path)
}
