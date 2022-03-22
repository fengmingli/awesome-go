/**
 * @Author LFM
 * @Date 2021/12/3 11:22 上午
 * @Since V1
 */

package typeconversion

import "fmt"

type People interface {
	Do()
}

type Man struct {
	Name string
}

func (man Man) Do() {
	fmt.Println(man.Name)
}

type Woman struct {
	Name string
}

func (woman *Woman) Do() {
	fmt.Println(woman.Name)
}

func SetMain() {
	//Man 值
	var man People = Man{Name: "男人"}
	man.Do()
	//
	var man2 People = &Man{Name: "男人2"}
	man2.Do()

	//Woman 引用
	var woman People = &Woman{}
	woman.Do()

	//var Woman2 People = Woman{}
	//Woman2.Do()
}
