package main

import "fmt"

func main() {
	/**
	   程序结构： 顺序结构、选择结构
	   选择结构: if分支语句、switch语句
	   switch ： 英文单词开关， 引申为选择
	   case：情况，事例
	   default：默认
	   灯的开关：只有两种状态，开/关
	   语法结构:
	     switch 变量 {
			case 值1 :
	        case 值2 :
	        case 值3 :
	        ....
	        default
	     }
	 */

	var age int
	// 18周岁 ：成年, 20周岁,女生法定结婚年龄，22周岁男生法定结婚年龄
	//
	fmt.Println("请输入你的年龄:")
	fmt.Scanln(&age)// 10
	switch age {
	case 18 :
		fmt.Println("恭喜你，成年了")
	case 20 :
		fmt.Println("如果你是女生，恭喜你达到法定结婚年龄了；如果你是男生，对不起还要再等等")
	case 22 :
		fmt.Println("不管你是男生还是女生，都达到了法定结婚年龄")
	default:
		fmt.Printf("你的年龄是:%d\n",age)
	}
}
