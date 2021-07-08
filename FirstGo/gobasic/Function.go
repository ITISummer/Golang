package main

import (
	"fmt"
)

func main() {
	c := MyFun(1,2)
	fmt.Println("函数测试，输出结果为：",c)

	d := fibo(8) // 1,1,2,3,5,8,13...
	fmt.Println("斐波那契数列求解结果为：",d)
	e := fibo2(8) // 1,1,2,3,5,8,13...
	fmt.Println("斐波那契数列求解结果为：",e)

	f := monkey(1)
	fmt.Println("猴子吃桃子问题：",f) // 1534
}

// 自定义函数
func MyFun(a int, b int) int {
	return a+b
}

// 递归求解斐波那契数列
func fibo(n int) int {
	if n == 1 || n ==2 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
// 递推求解斐波那契数列
func fibo2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	var a int = 1
	var b int = 1
	var sum int = 0
	for i := 1;i<=n-2;i++ {
		sum = a+b
		a = b
		b = sum
	}
	return sum
}

// 猴子吃桃子问题
func monkey(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的数值不对！")
		return 0
	}
	if n == 10 {
		return 1
	}
	return (monkey(n+1) + 1) * 2

}