package main

import (
	"container/list"
	"fmt"
)

func main() {
	var mylist list.List

	mylist.PushBack("go")
	mylist.PushBack("python")
	mylist.PushFront("java")

	// fmt.Println(mylist) // {{0xc000116540 0xc000116570 <nil> <nil>} 2}

	// Front 从头开始遍历，Next 往后找
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 反向遍历
	for i := mylist.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}

	// 插入元素
	i := mylist.Front()
	fmt.Println(i, "-----------")
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "go" {
			break
		}
	}
	// mylist.InsertBefore("TypeScript", i)
	mylist.Remove(i) // 删除元素

}
