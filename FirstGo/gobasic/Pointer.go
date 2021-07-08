package main
import "fmt"

func main() {
	// 普通类型
	var i int = 10
	fmt.Println("i 的地址是：",&i) // 0xc0000ac058
	
	// 指针类型
	var ptr *int = &i
	fmt.Println("ptr 指向的地址是：", ptr) //0xc0000ac058

	// 取指针指向的值
	fmt.Println("ptr 指向的地址中的值是：",*ptr)


}