package main

var variableB = "G"

func main() {
	// 输出结果为：GOO
	n2()
	m2()
	n2()
}

func n2() {
	print(variableB)
}

func m2() {
	// 直接赋值会改变变量本身
	variableB = "O"
	print(variableB)
}