/**
 * @Author LFM
 * @Date 2021/9/20 8:35 下午
 * @Since V1
 */

package chain

//HandlerFunc 责任链处理函数
type HandlerFunc func([]interface{}) ([]interface{}, bool)

//ResultHandler 放到最后，处理链的结果
type ResultHandler struct {
	Params string
}

func (f ResultHandler) ResultHandlerFunc(content []interface{}) ([]interface{}, bool) {
	return content, true
}

//AssemblyChainByFunc 组装责任链
func AssemblyChainByFunc(context []interface{}, functions ...HandlerFunc) []interface{} {
	functions = append(functions, ResultHandler{}.ResultHandlerFunc)
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
