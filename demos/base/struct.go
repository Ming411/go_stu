package main

import "fmt"

/* 结构体 */
type Person struct {
	name    string
	age     int
	address string
	height  float32
}

// 结构体嵌套
type Student struct {
	// p     Person  // 方式一  xxx.p.name
	Person // 方式二  匿名嵌套   xxx.name
	score  float32
}
type Teacher struct {
	Person struct {
		name string
		age  int
	}
}

// =========> 结构体定义方法
func (p Person) speak() {
	fmt.Println(p.name, "在说话")
}
func main() {

	p1 := Person{
		"张三",
		18,
		"北京",
		1.88,
	}
	// p2 := Person{
	// 	name:    "李四",
	// 	age:     20,
	// 	address: "上海",
	// 	height:  1.88,
	// }
	p1.speak() // 有点类的概念

	var persons []Person
	persons = append(persons, p1)

	persons2 := []Person{
		{
			name: "张三",
			age:  18,
		}, {
			name: "李四",
		},
	}
	fmt.Println("persons2:", persons2)

	// =====>赋值方式
	var p Person
	p.age = 20
	fmt.Println(p.height) // 0

	// =====> 匿名结构体
	address := struct {
		province string
		city     string
	}{
		"北京",
		"海淀",
	}
	fmt.Println(address.province)
}
