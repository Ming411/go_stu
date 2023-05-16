package main

import "fmt"

/* 指针 */
type Person struct {
	name string
}

func changeName(p *Person) {
	// (p Person)  值传递，默认会先拷贝一份，并不会修改原来的值
	// (p *Person) *表示指针 需要在传递时搭配取址符号 &p，会修改原来的值
	p.name = "coder"
}

// 接收者 指针传递
func (p *Person) speak() {
	p.name = "coder2"
}
func changeSlice(s []int) {
	// 切片是引用传递
	s[0] = 100
}

// 使用 * 运算符可以获取指针指向的变量的值，
// 而使用 & 运算符可以获取变量的地址，生成指向变量的指针
// 通过指针交换两个值，并非交换二者地址，交换的是二者地址内的值
func swap(a, b *int) {
	// a, b = b, a  // 由于是值传递，这里是复制的值
	t := *a // 临时值
	*a = *b
	*b = t
	// *a, *b = *b, *a

}

func main() {
	// 希望结构体传值的时候在函数中修改的值可以反应到变量中
	var p = Person{
		name: "ccc",
	}
	changeName(&p)
	// p.speak()
	fmt.Println(p)

	var s = []int{1, 2, 3}
	changeSlice(s)
	fmt.Println(s)

	// %p 是一个格式化输出占位符，用于将一个指针变量的值输出为十六进制表示的内存地址。
	var pi *Person = &p
	fmt.Printf("%p\n", pi)

	po := &Person{
		name: "aaa",
	}
	(*po).name = "bbb"
	// po.name = "bbb" // 效果相同
	fmt.Println(po)

	// ====================>
	// 指针必须初始化，否则会出现nil
	// --> 方式一
	// var aa Person
	// bb := &aa
	// fmt.Println(bb)
	// --> 方式二
	// p1 := &Person{}
	// --> 方式三 推荐使用
	var p2 = new(Person)
	fmt.Println("p2", p2)

	// ====================>
	c1 := 1
	c2 := 2
	swap(&c1, &c2)
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
}
