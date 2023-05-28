/**
 * @Author LFM
 * @Date 2021/9/20 8:35 下午
 * @Since V1
 */

package chain

//Handler 责任链处理函数
type Handler func(interface{}) (interface{}, bool)

//GetResultByChain 获取处理后的值
func GetResultByChain(key interface{}, functions ...Handler) interface{} {
	for _, f := range functions {
		value, ok := f(key)
		if ok {
			return value
		}
	}
	return nil
}
