package main

type Person struct {
}

func main() {
	/*
		不同类型的零值不同
		bool false
		numbers 0
		string ""
		pointer,slice,map,channel,interface nil
		struct 不是nil
		所以 一般都推荐声明即初始化 make()
	*/

	// map 和 slice 结果类似
	// var p1 []Person  // nil slice
	var p2 = make([]Person, 0) // empty slice,not nil
	if p2 == nil {
		println("p2 is nil")
	}
}
