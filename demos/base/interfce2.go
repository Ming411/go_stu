package main

import "fmt"

type Writer interface {
	Writer()
}
type MyReadWriter interface {
	Writer // 接口的继承
	ReadWrite()
}

/* interface{}  ~= any */
func add(a, b interface{}) interface{} {
	// 可能会转换失败，所以需要捕获error
	// ai, _ := a.(int) // a.()  类型断言

	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		return nil
	}
}

// ======>
func mPrint(data ...interface{}) {
	for _, value := range data {
		fmt.Println(value)
	}
}
func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))

	mPrint([]interface{}{1, 2, 3}...)
}
