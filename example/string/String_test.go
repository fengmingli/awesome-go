/**
 * @Author LFM
 * @Date 2021/11/23 4:46 下午
 * @Since V1
 */

package string

import (
	"fmt"
	"reflect"
	"testing"
)

type name struct {
	Name string
}

func TestString(t *testing.T) {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it) //3 int
	A()
}
func A() {
	xx := name{Name: "lfm"}
	xxx := name{Name: "lfm"}
	B(xx)
	fmt.Println(xx)
	C(&xxx)
	fmt.Println(xxx)
}
func B(name2 name) {
	name2.Name = "lfm-1"
}
func C(name2 *name) {
	name2.Name = "lfm-2"
}
