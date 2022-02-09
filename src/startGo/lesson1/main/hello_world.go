package main // 包，表明代码所在的模块，程序入口一定是main包，文件名可以不是main

import (
	"fmt"
	"os"
) // 引入代码以依赖

// 功能实现
func main() { //main方法不支持 传入参数和返回数据
	// 通过os的获取传入参数， 通过os返回数据
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
	}
	// 异常退出
	os.Exit(-1)
}
