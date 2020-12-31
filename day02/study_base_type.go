package main

func main()  {
	println(true)
	println(false)
	println(!true)
	println(!false)

	// 整数
	// 有符号整数
	// int8 的范围 -128～127，超出会产生编译错误：constant 128 overflows int8
	var a int8 = 127
	println(a)

	// -32768~32767
	var b int16 = 23232
	println(b)

	// -2147483648~2147483647
	var c int32 = 2147483647
	println(c)

	// -9,223,372,036,854,775,808~9,223,372,036,854,775,807
	var d int64 = 234234234234234234
	println(d)

	var e int = 922337203685477580
	println(e)

	// 无符号整数
	// 0~255
	var f uint8 = 255
	println(f)

	// 0~65,535
	var g uint16 = 65535
	println(g)

	// 0~4,294,967,295
	var h uint32 = 4294967295
	println(h)

	// 0~18,446,744,073,709,551,615
	var i uint64 = 18446744073709551615
	println(i)

	// +- 1e-45 -> +- 3.4 * 1e38
	var j float32 = 3
	println(j)

	// +- 5 * 1e-324 -> 107 * 1e308
	var k float64 = 3
	println(k)
}
