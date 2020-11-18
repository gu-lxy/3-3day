package main

import "fmt"

func main() {
	// 一年四季的例子
	// 1 => 第一季度
	// 2 => 第二季度
	var num int
	fmt.Println("请输入1-4之间的一个数值：")
	fmt.Scanln(&num)// 6
	switch num {// 6
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:// case 2 出现错误：Duplicate case 2
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	}
	fmt.Println("程序执行结束")
}
