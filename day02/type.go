package main

type ZT int
type Repe string

func main() {
	var a, b ZT = 3, 4
	c := a + b
	println(c)

	var strB Repe = "hello"
	println(strB)

	// byte 类型其实就是 uint8 的别名
	var ch byte = 'A'
	println(ch)

	var (
		t = 3
		bb = 5
	)

	println(t, bb)
	println(`this is a row string` == "this is a row string")
	print("this is a row string \n")
	println(len("go"))

	str := "abcd"
	println(&str)
}
