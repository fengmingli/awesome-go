package analysis

import "fmt"

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:12
 * @Since: 1.0.0
 * @Desc: TODO
 */

type A struct {
}

func NewA() Worker {
	return A{}

}

func (a A) Split() {
	fmt.Println("A-Split")
}

func (a A) Send() {
	fmt.Println("A-Send")
}
