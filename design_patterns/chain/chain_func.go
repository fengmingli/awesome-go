/**
 * @Author LFM
 * @Date 2021/9/20 8:35 下午
 * @Since V1
 */

package chain

import "strings"

type HandlerFunc func(string) (string, bool)

func AdHandlerFunc(content string) (string, bool) {
	newContent := strings.Replace(content, "广告", "**", 1)
	// try get value from memory cache
	return newContent, false
}

func YellowHandlerFunc(content string) (string, bool) {
	newContent := strings.Replace(content, "涉黄", "**", 1)
	// try get value from redis
	return newContent, false
}

func SensitiveHandlerFunc(content string) (string, bool) {
	newContent := strings.Replace(content, "敏感词", "***", 1)
	// try get value from mysql database
	return newContent, false
}

//FinallyHandlerFunc 放到最后，结束链的处理
func FinallyHandlerFunc(content string) (string, bool) {
	return content, true
}

func GetValueByChain(context string, functions ...HandlerFunc) string {
	functions = append(functions, FinallyHandlerFunc)
	var newContext = context
	for _, f := range functions {
		value, ok := f(newContext)
		newContext = value
		if ok {
			return value
		}
	}
	return context
}
