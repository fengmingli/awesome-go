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
	result := AssemblyChainByFunc([]interface{}{"我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容"},
		AdHandler{params: "广告"}.AdHandlerFunc,
		YellowHandler{params: "涉黄"}.YellowHandlerFunc)

	fmt.Println(result)
}
