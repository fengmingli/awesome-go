package main

import "fmt"

func modifyValueByPointer(ptr *int) {
	*ptr = 20
}

func modifyValueByReference(ref int) {
	ref = 30
}

func main() {
	var x int = 10

	// 使用指针修改变量的值
	fmt.Println("Before modification (pointer):", x)
	modifyValueByPointer(&x)
	fmt.Println("After modification (pointer):", x)

	// 使用引用修改变量的值
	fmt.Println("Before modification (reference):", x)
	modifyValueByReference(x)
	fmt.Println("After modification (reference):", x)

	ref := 10

	fmt.Println("Before modification (reference):", &ref)
}
