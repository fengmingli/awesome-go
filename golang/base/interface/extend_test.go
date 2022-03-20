/**
 * @Author LFM
 * @Date 2021/12/9 3:12 下午
 * @Since V1
 */

package _interface

import "testing"

func TestB(t *testing.T) {
	NewBDynamicClient().
		Resource().
		Namespace("xxxx").
		Create()

}
