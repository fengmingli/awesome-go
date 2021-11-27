/**
 * @Author LFM
 * @Date 2021/10/26 2:57 下午
 * @Since V1
 */

package copy

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	c := make([]int, 0)
	for i := 0; i < 10; i++ {
		c = append(c, i)
	}

	Copy(c, 5)

}

func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
