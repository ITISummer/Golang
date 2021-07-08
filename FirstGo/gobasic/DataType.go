package main

import "fmt"
import "unsafe"
import "strconv"

func main() {
	// 整数类型
	var v1 int = 1
	var v2 int8 = 127
	var v3 int16 = 128
	var v4 int32 = 129
	var v5 int64 = 130
	fmt.Println("v1 = ",v1,"\nv2 = ",v2,"\nv3 = ",v3,"\nv4 = ",v4,"\nv5 = ",v5)

	var v6 uint8 = 255
	fmt.Println("v6 = ",v6)

	fmt.Printf("v6's data type is %T | v6's byte size is %d\n",v6, unsafe.Sizeof(v6))

	fmt.Println("----------------------------------")
	// 浮点类型
	var price float32 = 89.12
	fmt.Println("price = ", price)
	var num1 float32 = -0.00089
	var num2 float64 = -78909656.09
	fmt.Println("num1 = ", num1,"num2 = ",num2)

	// 尾数部分可能丢失，造成精度损失。
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 = ", num3,"num4 = ",num4)

	// Golang 的浮点类型默认声明为 float64 类型
	var num5 = 1.1
	fmt.Printf("num5 的数据类型为 %T\n",num5)


	fmt.Println("----------------------------------")
	// 字符类型
	var c1 byte = 'a'
	var c2 byte = '0' // 字符 0

	// 当我们直接输出 byte 值，就是输出了对应的字符的码值
	fmt.Println("c1 = ",c1)
	fmt.Println("c2 = ",c2)
	// 如果希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1 = %c c2 = %c\n",c1,c2)

	// var c3 byte = '吕' // overflow 溢出
	var c3 int = '吕'
	fmt.Printf("c3 = %c c3 对应码值= %d\n",c3,c3)

	fmt.Println("----------------------------------")

	// bool 类型
	b := true
	fmt.Printf("b 的类型是：%T, b 的大小是：%d \n",b, unsafe.Sizeof(b))


	fmt.Println("----------------------------------")
	// string 类型
	var address string = "北京,lvlv"
	fmt.Println(address)

	fmt.Println("----------------------------------")
	// 基本数据类型转 string 类型 - 第一种方式 Sprintf()
	var str string
	str = fmt.Sprintf("%d",99)
	fmt.Printf("整型转字符串：%T -- %q\n",str,str)
	str = fmt.Sprintf("%f",99.99)
	fmt.Printf("浮点转字符串：%T -- %q\n",str,str)
	str = fmt.Sprintf("%t",true)
	fmt.Printf("bool转字符串：%T -- %q\n",str,str)
	str = fmt.Sprintf("%c",'h')
	fmt.Printf("字符转字符串：%T -- %q\n",str,str)
	// 基本数据类型转 string 类型 - 第二种方式 strconv
	fmt.Println("----------------------------------")
	str = strconv.FormatInt(32,10)
	fmt.Printf("整型转字符串：%T -- %q\n",str,str)
	// 'f' 表示格式，10 表示小数位保留 10 位，64 表示这个小数是 float64
	str = strconv.FormatFloat(99.999,'f',10,64)
	fmt.Printf("浮点转字符串：%T -- %q\n",str,str)
	str = strconv.FormatBool(false)
	fmt.Printf("布尔转字符串：%T -- %q\n",str,str)
	fmt.Println("----------------------------------")
	var boo bool
	// boo,_ = strconv.ParseBool(str)
	boo,_ = strconv.ParseBool("true")
	fmt.Printf("字符串转布尔：%T -- %v\n",boo,boo)

}
