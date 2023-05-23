package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	// panic 会直接导致程序退出
	// 一般用于服务是启动失败，抛出错误
	// panic("this is a panic")
	return 0, errors.New("error")
}

func main() {
	// 这种简写方式 变量是在外面拿不到的
	if _, err := A(); err != nil {
		fmt.Println(err)
	}

	// recover会捕获panic错误
	// 必须要放在defer函数中
	// 后定义的defer会先执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到了错误", r)
		}
	}()
	var names map[string]string
	names["ccc"] = "coder"
}
