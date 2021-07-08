package main

import "fmt"

func main() {
	// exec 1
	var count int = 0;
	var sum int = 0;
	for i := 1;i<=100;i++ {
		if i % 9 == 0 {
			count++;
			sum += i;
		}
	}
	fmt.Println("1~100 是 9 的倍数的个数为：",count)
	fmt.Println("1~100 是 9 的倍数的的数和：",sum)

	/*
	// exec 2
	// 输入一个数，然后输出相应表达式
	var n int;
	fmt.Scanln(&n)
	for i := 0;i <= n; i++ {
		fmt.Printf("%v + %v = %v\n",i,n-i,n)
	}
	*/

	// exec 3
	// 打印金字塔案例
	var n int;
	fmt.Scanln(&n)
	for i:=1;i<=n;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	// exec 4
	// 打印金字塔案例
	var m int;
	fmt.Scanln(&m)
	for i:=1;i<=m;i++ {
		for k:=1;k<=m-i;k++{
			fmt.Printf(" ")
		}
		for j:=1;j<=2*i-1;j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}