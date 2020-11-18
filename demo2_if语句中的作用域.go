package main

import "fmt"

func main() {
	a := 10
	if a > 0 {
		fmt.Println(" a 是一个正数")
	}

	//go语言支持合并写法
	// if 条件表达式  表达式的值只有两种：true 、false
	if b := 6; b > 0 {
		fmt.Printf("b的值是：%d\n",b)
		fmt.Println("b是一个正数")
	}
	// b变量的作用范围是13行到16行
	//fmt.Printf("b的值是：%d\n",b)// 问题：b变量为什么是红色，程序会报错

	if a > 5 {
		c := 80
		fmt.Println(c)//  c变量的作用域是21行和22行。
	}
	//fmt.Println(c)// 超出了c变量的作用范围


}
