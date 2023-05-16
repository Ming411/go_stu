package main

import "fmt"

func main() {
	sliceList := []int{1, 2, 3, 4, 5}

	fmt.Println(sliceList[1:3]) // [2 3]
	fmt.Println(sliceList[1:])  // [2 3 4 5]
	fmt.Println(sliceList[:3])  //  [1 2 3]
	fmt.Println(sliceList[:])   // 就相当于复制了一份

	sliceList = append(sliceList, 6, 7) // 追加元素
	fmt.Println(sliceList)

	// =========> 合并切片
	courses1 := []string{"golang", "java", "python"}
	courses2 := []string{"golang1", "java1", "python"}
	courses1 = append(courses1, courses2...)
	// 并不会进行去重操作
	// [golang java python golang1 java1 python]
	fmt.Println(courses1)

	// =========> 浅拷贝slice
	// coursesCopy := courses2
	coursesCopy := courses2[:]
	// courses2[0] = "golang2" // 复制的仅仅是地址而已
	fmt.Println(coursesCopy)

	// ========> 深拷贝
	// make 可以预先给切片分配空间，如果不进行空间分配
	// 虽然是拷贝内容，但是没有空间拷贝不过来
	var coursesCopy2 = make([]string, len(courses2))
	copy(coursesCopy2, courses2)
	// courses2[0] = "golang2"
	fmt.Println(coursesCopy2)

	// ==============>  内部是值传递
}
