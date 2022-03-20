package _chan

import (
	"fmt"
	"time"
)

func NewInput() *input {
	return &input{
		input1: make(chan int),
		input2: make(chan int),
		output: make(chan int),
	}
}

type input struct {
	input1 chan int
	input2 chan int
	output chan int
}

func (i *input) MergeInput() {
	go i.merge(i.input1, i.input2, i.output)
	go i.in1()
	go i.in2()
	go i.out()
	time.Sleep(time.Second * 5)
	fmt.Println("主线程退出")
}

func (in *input) merge(in1, in2, out chan int) {
	for {
		select {
		case v := <-in1:
			out <- v
		case v := <-in2:
			out <- v
		}
	}
}

func (in *input) in1() {
	for i := 0; i < 10; i++ {
		in.input1 <- i
	}
	close(in.input1)
}

func (in *input) in2() {
	for i := 20; i < 30; i++ {
		in.input2 <- i
	}
	close(in.input2)
}

func (in *input) out() {
	for {
		select {
		case value := <-in.output:
			fmt.Println("输出：", value)
		}
	}
}

// TestQuit 测试channel用于通知中断退出的问题
func TestQuit() {
	g := make(chan int)
	quit := make(chan bool)
	mainQuit := make(chan bool)
	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				mainQuit <- true
			}
		}
	}()

	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	<-mainQuit
	fmt.Println("testAB退出")
}

func TestPCB() {
	fmt.Println("TestPCB")
	intChan := make(chan int)
	quitChan1 := make(chan bool)
	quitChan2 := make(chan bool)
	value := 0

	go func() {
		for i := 0; i < 3; i++ {
			value += 1
			intChan <- value
			fmt.Println("Write finish,value=", value)
			time.Sleep(time.Second)
		}
		quitChan1 <- true
	}()

	go func() {
		for {
			select {
			case v := <-intChan:
				fmt.Println("read finish, value ", v)
			case <-quitChan1:
				quitChan2 <- true
				return
			}
		}
	}()

	<-quitChan2
	fmt.Println("task is done ")
}

func NewTestPCB() {
	TestPCB()
}
