package main

/**
 * @Author: LFM
 * @Date: 2023/5/28 23:17
 * @Since: 1.0.0
 * @Desc: TODO
 */

import (
	"fmt"
)

// 策略接口
type PaymentStrategy interface {
	Pay(amount float64)
}

// 具体策略：支付宝支付
type AlipayStrategy struct{}

func (a *AlipayStrategy) Pay(amount float64) {
	fmt.Printf("支付宝支付成功，金额：%.2f\n", amount)
}

// 具体策略：微信支付
type WeChatPayStrategy struct{}

func (w *WeChatPayStrategy) Pay(amount float64) {
	fmt.Printf("微信支付成功，金额：%.2f\n", amount)
}

// 工厂模式：支付策略工厂
type PaymentStrategyFactory struct{}

func (f *PaymentStrategyFactory) CreatePaymentStrategy(paymentMethod string) PaymentStrategy {
	switch paymentMethod {
	case "alipay":
		return &AlipayStrategy{}
	case "wechat":
		return &WeChatPayStrategy{}
	default:
		panic("Unsupported payment method")
	}
}

func main() {
	// 创建支付策略工厂
	factory := &PaymentStrategyFactory{}

	// 使用支付宝支付
	alipayStrategy := factory.CreatePaymentStrategy("alipay")
	alipayStrategy.Pay(100.00)

	// 使用微信支付
	wechatPayStrategy := factory.CreatePaymentStrategy("wechat")
	wechatPayStrategy.Pay(200.00)
}
