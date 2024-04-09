package analysis

import "fmt"

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:12
 * @Since: 1.0.0
 * @Desc: TODO
 */

type B struct {
}

func NewB() Worker {
	return B{}

}

func (a B) Split() {
	fmt.Println("B-Split")
}

func (a B) Send() {
	fmt.Println("B-Send")
}
