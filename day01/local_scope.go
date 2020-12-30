package main

var variableA = "G"

func main() {
	// 输出结果为：GOG
	n()
	m()
	n()
}

func n() {
	print(variableA)
}

// 你可以在某个代码块的内层代码块中使用相同名称的变量，则此时外部的同名变量将会暂时隐藏
//（结束内部代码块的执行后隐藏的外部同名变量又会出现，而内部同名变量则被释放），你任何的操作都只会影响内部代码块的局部变量
func m() {
	variableA := "O"
	print(variableA)
}
