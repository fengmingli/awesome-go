/**
 * @Author LFM
 * @Date 2021/9/20 8:00 下午
 * @Since V1
 */

package chain

import (
	"fmt"
	"strings"
)

//AdHandler 广告过滤
type AdHandler struct {
	params string
}

func (ad AdHandler) AdHandlerFunc(content []interface{}) ([]interface{}, bool) {
	fmt.Println("执行广告过滤。。。")
	newContent := strings.Replace(content[0].(string), ad.params, "**", 1)
	fmt.Println(newContent)
	return []interface{}{newContent}, false
}

//YellowHandler 涉黄过滤
type YellowHandler struct {
	params string
}

func (yellow YellowHandler) YellowHandlerFunc(content []interface{}) ([]interface{}, bool) {
	fmt.Println("执行涉黄过滤。。。")
	newContent := strings.Replace(content[0].(string), yellow.params, "**", 1)
	fmt.Println(newContent)
	return []interface{}{newContent}, false
}
