package main

import "fmt"

func main() {
	// num1, num2 :=  3,5
	// 数学运算： 加减乘除
	fmt.Println("请输入两个整数:")
	var num1 int
	var num2 int
	fmt.Scanln(&num1,&num2)

	//输入想要执行的操作
	fmt.Println("请输入要执行的操作符号(+、-、*、/):")
	var str string
	fmt.Scanln(&str) // + 、 - 、* 、/

	switch str { // + 、 - 、* 、/
		case "+":
			fmt.Println("两数相加的结果是:",num1 + num2)
		case "-":
			fmt.Println("两数相减的结果是:",num1 -num2)
		case "*":
			fmt.Println("两数相乘的结果是:",num1 * num2)
		case "/":
			fmt.Println("两数相除的结果是：",num1/num2)
		default:
			fmt.Println("对不起，我暂时不能提供您想要的操作")
	}
}
