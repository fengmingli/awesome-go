/**
 * @Author LFM
 * @Date 2021/10/14 8:10 下午
 * @Since V1
 */

package deal_lock

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestService(t *testing.T) {
	//让pprof服务运行起来
	go StartHTTPDebug()

	//模拟两个处理请求的goroutine
	go worker1()
	go worker2()

	//下面3行只是为了让主协程不退出
	fmt.Println("input anything to quit.")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
