package main

import (
	"fmt"
	"tool/args"
)

// 注册命令行服务

// 初始化

// 预处理命令行解析
func main() {
	a := &args.Arguments{}
	err := a.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
}
