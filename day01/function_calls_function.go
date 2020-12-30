package main

var strA string

func main() {
	// 输出结果为：GOG
	strA = "G"
	print(strA)
	f1()
}

func f1() {
	strA := "O"
	print(strA)
	f2()
}

func f2() {
	print(strA)
}
