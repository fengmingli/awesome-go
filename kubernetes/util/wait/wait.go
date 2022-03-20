package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

func main() {

	go func1()
	go func2()
}

func func1() {
	stopCh := make(chan struct{})
	//初始化一个计数器
	i := 0
	wait.Until(func() {
		fmt.Printf("----%d----\n", i)
		i++
	}, time.Second*1, stopCh)
}

func func2() {
	stopCh := make(chan struct{})
	//初始化一个计数器
	i := 0
	go wait.Until(func() {
		fmt.Printf("----%d----\n", i)
		i++
	}, time.Second*1, stopCh)

	time.Sleep(time.Second * 10)
	stopCh <- struct{}{}

	fmt.Println("---上面的go routines 结束----")

	// 主程序，再休息3秒钟，再结束
	time.Sleep(time.Second * 3)

	fmt.Println("---主程序结束----")
}
