package main

import "fmt"

func main() {
	// 用户登录: 教辅平台,   131....2341   密码：20200303
	var  username  string // 用户名
	var password string // 用户的密码
	fmt.Println("请输入用户名和密码")
	fmt.Println("用户名:")
	fmt.Scanln(&username)// 读取键盘输入并复制
	fmt.Println("密码：")
	fmt.Scanln(&password)

	fmt.Printf("您输入的用户名是:%s,密码是:%s\n",username,password)

	var num int // 接受int类型
	fmt.Scanln(&num)
	if num == 1{
		fmt.Println("星期一")
	}else if num == 2{
		fmt.Println("星期二")
	}else if num == 3{
		fmt.Println("星期三")
	}else if num == 4{
		fmt.Println("星期四")
	}else if num == 5{
		fmt.Println("星期五")
	}else if num == 6{
		fmt.Println("星期六")
	}else if num == 7{
		fmt.Println("星期日")
	}else {
		fmt.Println("错误信息")
	}
}
