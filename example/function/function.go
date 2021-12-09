/**
 * @Author LFM
 * @Date 2021/12/3 10:58 下午
 * @Since V1
 */

package function

import (
	"fmt"
	"time"
)

func function() {
	// false, print 3 3 3
	values := []int{1, 2, 3}
	for _, val := range values {
		go func() {
			fmt.Println("==1==", val)
		}()
	}
	// true, print 1 2 3
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println("==2==", val)
		}(val)
	}
	time.Sleep(10 * time.Second)
}

func function2() {
	j := 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
