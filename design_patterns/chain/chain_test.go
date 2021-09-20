/**
 * @Author LFM
 * @Date 2021/9/20 8:03 下午
 * @Since V1
 */

package chain

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}
	// 将责任链串起来
	adHandler.handler = yellowHandler
	yellowHandler.handler = sensitiveHandler

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}

func TestChainFunc(t *testing.T) {
	fmt.Println(GetValueByChain("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容",
		AdHandlerFunc,
		SensitiveHandlerFunc,
		YellowHandlerFunc))
}
