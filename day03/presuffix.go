package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "this is an example"

	
	fmt.Println(str)
	
	// 是否是以 this 为前缀
	fmt.Println(strings.HasPrefix(str, "this"))
	// 是否以 examples 为后缀
	fmt.Println(strings.HasSuffix(str, "examples"))
	// 获取字符串 "i" 在字符串中第一次出现的索引，-1 表示没有找到(不包含该字符串)
	fmt.Println(strings.Index(str, "i"))
	// 获取字符串 'i" 在字符串中最后一次出现的索引，-1 表示没有找到
	fmt.Println(strings.LastIndex(str, "i"))
}
