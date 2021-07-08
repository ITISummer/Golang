package main
import "fmt"

func main() {
	// 指定变量类型
	var i int
	i = 10
	fmt.Println("i = ",i)
	// 根据值自行判定变量类型(类型推导)
	var str = "一个字符串"
	fmt.Println(str)
	// 省略 var，注意：:= 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// 下面的方式等价于 var name string name = "lvlv"
	name := "lvlv"
	fmt.Println(name)
	// 一次性声明多个变量 - 方式一
	var v1,v2,v3 int
	v1 = 1
	v2 = 2
	v3 = 3
	fmt.Println(v1,v2,v3)
	// 一次性声明多个变量 - 方式二
	var num,nam,str2 = 100,"lvlv","str"
	fmt.Println(num,nam,str2)
	// 一次性声明多个变量 - 方式三 - 使用类型推导
	n1,n2,n3 := 100,"lvlv","str"
	fmt.Println(n1,n2,n3)
	// 一次性声明多个变量 - 方式四 - 声明全局变量
	var(
		n4 = 300
		n5 = 900
		n6 = "mary"
	)
	fmt.Println(n4,n5,n6)
	// 字符串拼接
	fmt.Println(nam+" "+n6)

}