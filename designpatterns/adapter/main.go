package main

import "fmt"

/**
 * 适配器模式
 * @Author: LFM
 * @Date: 2023/5/28 00:06
 * @Since: 1.0.0
 * @Desc: TODO
 */

// Target 目标接口
type Target interface {
	Request()
}

// AdapterFunc 适配者接口
type AdapterFunc interface {
	SpecificRequest()
}

// ConcreteAdapterFunc 具体的适配者实现
type ConcreteAdapterFunc struct{}

func (a *ConcreteAdapterFunc) SpecificRequest() {
	fmt.Println("Specific Request")
}

// Adapter 适配器结构体
type Adapter struct {
	AdapterFunc
}

func (a *Adapter) Request() {
	// 调用适配者的方法
	a.SpecificRequest()
}

// 客户端代码
func main() {
	// 创建适配者对象
	adapterFunc := &ConcreteAdapterFunc{}
	// 创建适配器对象，将适配者对象传入
	adapter := &Adapter{AdapterFunc: adapterFunc}
	// 调用适配器的目标方法
	adapter.Request() // 输出 "Specific Request"
}
