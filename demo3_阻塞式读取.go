package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)// 键盘上输入18  阻塞式: 一直等待输入，直到用户敲入回车
	fmt.Printf("你的年龄是:%d\n",age)// 第9行不会主动执行
}
