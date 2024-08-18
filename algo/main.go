package main

import (
	"container/list"
	"fmt"
	"time"
)

/**
 * @Author: LFM
 * @Date: 2024/6/22 23:24
 * @Since: 1.0.0
 * @Desc: TODO
 */

func main() {
	start := time.Now()
	sum := forLoopRecur(100)
	end := time.Since(start)
	fmt.Println(sum)
	fmt.Printf("doSomething() took %s\n", end)

}

func add(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func dg(n int) int {
	if n == 1 {
		return 1
	}
	res := dg(n - 1)
	return n + res
}

/* 使用迭代模拟递归 */
func forLoopRecur(n int) int {
	// 使用一个显式的栈来模拟系统调用栈
	stack := list.New()
	res := 0
	// 递：递归调用
	for i := n; i > 0; i-- {
		// 通过“入栈操作”模拟“递”
		stack.PushBack(i)
	}
	// 归：返回结果
	for stack.Len() != 0 {
		// 通过“出栈操作”模拟“归”
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	// res = 1+2+3+...+n
	return res
}
